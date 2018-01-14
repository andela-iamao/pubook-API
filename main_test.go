package main

import (
	"net/http"
	"testing"
	. "github.com/emicklei/forest"
)


var client = NewClient("http://localhost:5000", new(http.Client))

func TestGetAllBooksSuccess(t *testing.T) {
	cfg := NewConfig("/api/v1/books").Header("Accept", "application/json")
	r := client.GET(t, cfg)
	ExpectStatus(t, r, 200)
}

func TestCreateBookSuccess(t *testing.T) {
	cfg := NewConfig("/api/v1/books").
		Header("Accept", "application/json").
			Content("{\"title\": \"ABC\", \"author\": \"Me\"}", "application/json")
	r := client.POST(t, cfg)
	ExpectStatus(t, r, 201)
}

func TestGetOneBookSuccess(t *testing.T) {
	TestCreateBookSuccess(t)

	getCfg := NewConfig("/api/v1/book/1").
		Header("Accept", "application/json")
	getRes := client.GET(t, getCfg)
	ExpectStatus(t, getRes, 200)
}

func TestUpdateBookSuccess(t *testing.T) {
	TestCreateBookSuccess(t)

	cfg := NewConfig("/api/v1/book/1").
		Header("Accept", "application/json").
		Content("{\"title\": \"EFG\"}", "application/json")
	r := client.PUT(t, cfg)
	ExpectStatus(t, r, 204)
}

func TestDeleteBookSuccess(t *testing.T) {
	TestCreateBookSuccess(t)

	cfg := NewConfig("/api/v1/book/1").
		Header("Accept", "application/json")
	r := client.DELETE(t, cfg)
	ExpectStatus(t, r, 204)
}