package gocdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// DocClient represents the client to the Cosmos DB
type DocClient struct {
	Endpoint *url.URL
	AuthKey  string
}

func (c *DocClient) GetDatabase(dbId string) (*Database, error) {
	ri := fmt.Sprintf("dbs/%s", dbId)
	url := fmt.Sprintf("%s/dbs/%s", c.Endpoint, dbId)
	utcDate := utcNow()
	authSig := generateAuthSig("Get", "dbs", ri, utcDate, c.AuthKey, "master", "1.0")

	req, _ := createRequest("GET", url, utcDate, authSig, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get database. Error: %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	defer resp.Body.Close()

	db := &Database{}

	if err := json.NewDecoder(resp.Body).Decode(db); err != nil {
		return nil, fmt.Errorf("JSON decode failed. Error: %v", err)
	}

	return db, nil
}

// func (c *DocClient) CreateDatabase(dbId string) (*Database, error) {

// }

func (c *DocClient) CreateDatabaseIfNotExist(dbId string) (*Database, error) {
	db, err := c.GetDatabase(dbId)
	if err != nil {
		return nil, err
	}

	if db != nil {
		return db, nil
	}

	return nil, nil
}
