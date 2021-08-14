package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

// The tests in this suite check if the business requirements

// Requirement 1: Creates a new basket by calling POST on /Baskets
func TestCreateNewBasket(t *testing.T) {

	//Initialize docker
	var db *sql.DB
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	//Run CockroachDB container
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Hostname:   "db",
		Name:       "db",
		Repository: "cockroachdb/cockroach",
		Tag:        "latest",
		Cmd:        []string{"start-single-node", "--insecure"}})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgresql://root@localhost:%s/events?sslmode=disable", resource.GetPort("26257/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to cockroach container: %s", err)
	}

	//Run our test app container
	resource2, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "lana-sre-rest",
		Tag:        "latest",
		Env:        []string{"PGHOST=db", "PGPORT=26257", "PGDATABASE=postgres", "PGUSER=root", "PGPASSWORD=secret"},
		Links:      []string{"db:db"}})
	require.NoError(t, err, "could not start container")

	//Create a new basket
	var resp *http.Response
	postBody, _ := json.Marshal("")

	err = pool.Retry(func() error {
		resp, err = http.Post(fmt.Sprint("http://localhost:", resource2.GetPort("3000/tcp"), "/Baskets"), "application/json", bytes.NewBuffer(postBody))
		if err != nil {
			t.Log(err)
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	require.Contains(t, string(body), "Basket created successfully")

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// When you're done, kill and remove the container
	if err = pool.Purge(resource2); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

}
