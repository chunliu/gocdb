package gocdb

import (
	"testing"
)

func TestCreateDocumentCollection(t *testing.T) {
	client := getDocClient()
	collection := &Collection{}
	collection.Id = "Collection1"

	var err error
	collection, err = client.CreateDocumentCollection("db", collection)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
}
