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

func TestReplaceDocumentCollection(t *testing.T) {
	client := getDocClient()
	collection := &Collection{}
	collection.Id = "Collection1"
	collection.IndexingPolicy = IndexingPolicy{
		IndexMode:     "Lazy",
		Automatic:     true,
		ExcludedPaths: []ExcludedPath{},
		IncludedPaths: []IncludedPath{
			IncludedPath{
				Path: "/*",
				Indexes: []Index{
					Index{
						DataType:  "Number",
						Precision: -1,
						Kind:      "Range",
					},
					Index{
						DataType:  "String",
						Precision: 3,
						Kind:      "Hash",
					},
				},
			},
		},
	}

	var err error
	collection, err = client.ReplaceDocumentCollection("db", collection)
	if err != nil {
		t.Errorf("Failed. %v", err)
	}
}
