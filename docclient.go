package gocdb

import (
	"net/url"
)

// DocClient represents the client to the Cosmos DB
type DocClient struct {
	Endpoint url.URL
	AuthKey  string
}
