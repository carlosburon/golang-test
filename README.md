# Lana SRE Challenge

## Description

This repository implements a simple checkout HTTP API that allows the following actions:

    - Create a new checkout basket
    - Add a product to a basket
    - Get the total amount in a basket
    - Remove the basket


## Requirements

- **Golang 1.16**: *For the app implementation*
- **Docker 3.7**: *For deploying the app*
- **Github Actions**: *For the CI pipeline*
- **CockroachDB**: *As postgreSQL databate*
- **Ory/Dockertest**: *For the unit tests*
- **Docker compose (optional)**: *For building and testing the application in the inner loop before pushing*

- **Network reqs**:
    - *Ports exposed:*
        - App http: 3000
        - Postgres: 26257
        - Postgres http: 8080

## Files

Some files from the source to highlight

- **📄 main.go** *The app implemented*

- **📄 .github/workflows/ci-pipeline.yml** *The CI pipeline*

- **🧪 main_test.go** *Unit tests*

- **🐋 Dockerfile** *For building the docker image*

- **🐋 docker-compose.yml** *Inner loop pipeline for testing in local*

- **📃 curl.md** *Curl examples for testing the endpoints*

- **📃 .env** *Recommended for setting up PGPASSWORD for future releases where the DB connection is securized*

## Usage

If you want to test building the app in local, run the following docker compose command
    - docker-compose up --build

A simple unit test can be run using
    - go test -v ./..

You can test the endpoints manually with the examples provided in the **📃 curl.md** file provided in this repository

