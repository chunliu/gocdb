package gocdb

import (
	"net/url"
	"testing"
)

func TestGetDatabase(t *testing.T) {
	endpoint := "https://localhost:8081"
	key := "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw=="
	client := &DocClient{}
	client.Endpoint, _ = url.Parse(endpoint)
	client.AuthKey = key

	client.GetDatabase("db")
}
