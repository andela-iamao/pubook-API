#!/bin/bash

export GIN_MODE=test
go run main.go &
go test

kill $(lsof -t -i:2100)