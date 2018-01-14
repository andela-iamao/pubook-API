#!/bin/bash

export GIN_MODE=test
export PUBOOK_DBNAME=pubook-db-test
export PUBOOK_HOST=127.0.0.1
export PUBOOK_USER=postgres
export PUBOOK_PASSWORD=

go run main.go &
go test

kill $(lsof -t -i:2100)