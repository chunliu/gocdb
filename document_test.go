package gocdb

import (
	"fmt"
	"testing"
)

type message struct {
	Id   string `json:"id"`
	Name string
	Body string
}

func TestCreateDocument(t *testing.T) {
	client := getDocClient()

	msg := &message{
		Id:   "testdoc123",
		Name: "Andersen",
		Body: "This is a test",
	}
	doc := &DocumentBase{}

	type document struct {
		*DocumentBase
		*message
	}

	c := &document{doc, msg}

	ret, err := client.CreateDocument("db", "Collection1", c, RequestOptions{})
	if err != nil {
		t.Errorf("Failed. %v, %v", err, ret)
	}

	c = ret.(*document)
	fmt.Println(c.Name)
}
