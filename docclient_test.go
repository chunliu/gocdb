package gocdb

import (
	"net/url"
	"testing"
)

func getDocClient() *DocumentClient {
	endpoint := "https://localhost:8081"
	key := "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw=="
	client := &DocumentClient{}
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
}

func TestCreateDatabase(t *testing.T) {
	client := getDocClient()

	_, err := client.CreateDatabase("db1")
	if err != nil && err.Error() != "CreateDatabase: HTTP Status: 409" {
		t.Errorf("Failed, %s", err)
	}
}

func TestCreateDatabaseIfNotExist(t *testing.T) {
	client := getDocClient()

	db, err := client.CreateDatabaseIfNotExist("db")
	if err != nil {
		t.Errorf("Failed, %s", err)
	}

	if db.Id != "db" {
		t.Errorf("Failed, %s", db.SelfLink)
	}

	db, err = client.CreateDatabaseIfNotExist("db1")
	if err != nil {
		t.Errorf("Failed, %s", err)
	}

	if db.Id != "db1" {
		t.Errorf("Failed, %s", db.SelfLink)
	}
}

func TestListDatabases(t *testing.T) {
	client := getDocClient()

	dbs, err := client.ListDatabases()
	if err != nil {
		t.Errorf("Failed. %v", err)
	}

	if dbs.Count <= 0 {
		t.Errorf("Failed. %d", dbs.Count)
	}
}

func TestDeleteDatabase(t *testing.T) {
	client := getDocClient()

	err := client.DeleteDatabase("db1")
	if err != nil {
		t.Errorf("Failed. %v", err)
	}
}
