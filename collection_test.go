package gocdb

import (
	"testing"
)

func TestCreateDocumentCollection(t *testing.T) {
	client := getDocClient()
	collection := &Collection{}
	collection.Id = "Collection1"

	var err error
	collection, err = client.CreateDocumentCollection("db", collection, RequestOptions{})
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
}

func TestGetDocumentCollection(t *testing.T) {
	client := getDocClient()

	collection, err := client.GetDocumentCollection("db", "Collection1")
	if err != nil {
		t.Errorf("Failed: %v", err)
	} else {
		t.Logf("%s", collection.DocumentsLink)
	}
}

func TestDeleteDocumentCollection(t *testing.T) {
	client := getDocClient()

	collection := &Collection{}
	collection.Id = "Collection2"

	collection, err := client.CreateDocumentCollectionIfNotExist("db", collection, RequestOptions{})
	if err != nil {
		t.Errorf("Failed. %v", err)
	}

	if collection.Id != "Collection2" {
		t.Errorf("Failed.")
	}

	err = client.DeleteDocumentCollection("db", "Collection2")
	if err != nil {
		t.Errorf("Failed. %v", err)
	}
}