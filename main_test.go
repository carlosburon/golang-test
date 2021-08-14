package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

// The tests in this suite check if the business requirements

// Requirement 1: Creates a new basket by calling POST on /Baskets
func TestCreateNewBasket(t *testing.T) {

	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	resource, err := pool.Run("lana-sre-rest", "latest", []string{})
	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})

	var resp *http.Response

	postBody, _ := json.Marshal("")

	err = pool.Retry(func() error {
		resp, err = http.Post(fmt.Sprint("http://localhost:", resource.GetPort("3000/tcp"), "/Baskets"), "application/json", bytes.NewBuffer(postBody))
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
}
