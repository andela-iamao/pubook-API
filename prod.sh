#!/bin/bash

export GIN_MODE=release

go get github.com/emicklei/forest
go get github.com/lib/pq
go get github.com/gin-gonic/gin
go get github.com/astaxie/beego/orm

go build -o bin/application main.go
