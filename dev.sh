#!/bin/bash

export GIN_MODE=debug
export PUBOOK_DBNAME=pubook-db-debug
export PUBOOK_HOST=127.0.0.1
export PUBOOK_USER=postgres
export PUBOOK_PASSWORD=

go run main.go