package gocdb

import (
	"net/url"
	"testing"
)

func getDocClient() *DocClient {
	endpoint := "https://localhost:8081"
	key := "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw=="
	client := &DocClient{}
	client.Endpoint, _ = url.Parse(endpoint)
	client.AuthKey = key

	return client
}
func TestGetDatabase(t *testing.T) {
	client := getDocClient()

	db, err := client.GetDatabase("db")
	if err != nil {
		t.Errorf("Failed, %s", err)
	}

	if db.Id != "db" {
		t.Errorf("Failed")
	}

	db, err = client.GetDatabase("db1")
	if db != nil {
		t.Errorf("Failed")
	}
}
