package gocdb

import (
	"bytes"
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

// GetDatabase restrieves the database info based on the database Id
func (c *DocClient) GetDatabase(dbID string) (*Database, error) {
	verb := "GET"
	ri := fmt.Sprintf("dbs/%s", dbID)
	url := fmt.Sprintf("%s/dbs/%s", c.Endpoint, dbID)
	utcDate := utcNow()
	authSig := generateAuthSig(verb, "dbs", ri, utcDate, c.AuthKey, "master", "1.0")

	req, err := createRequest(verb, url, utcDate, authSig, nil)
	if err != nil {
		return nil, fmt.Errorf("GetDatabase: Failed to create request. Error: %v", err)
	}

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetDatabase: Failed to get database. Error: %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	defer resp.Body.Close()

	db := &Database{}

	if err := json.NewDecoder(resp.Body).Decode(db); err != nil {
		return nil, fmt.Errorf("GetDatabase: JSON decode failed. Error: %v", err)
	}

	return db, nil
}

// CreateDatabase create a database with the Id
func (c *DocClient) CreateDatabase(dbID string) (*Database, error) {
	verb := "POST"
	url := fmt.Sprintf("%s/dbs", c.Endpoint)
	utcDate := utcNow()
	// ResourceId is empty in this post request
	authSig := generateAuthSig(verb, "dbs", "", utcDate, c.AuthKey, "master", "1.0")

	db := &Database{}
	db.Id = dbID

	jv, _ := json.Marshal(db)

	req, err := createRequest(verb, url, utcDate, authSig, bytes.NewBuffer(jv))
	if err != nil {
		return nil, fmt.Errorf("CreateDatabase: Failed to create request. Error: %v", err)
	}

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("CreateDatabase: Failed to create database. Error: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("CreateDatabase: Failed to create database. Error: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(db); err != nil {
		return nil, fmt.Errorf("CreateDatabase: JSON decode failed. Error: %v", err)
	}

	return db, nil
}

// CreateDatabaseIfNotExist retrieve the database if it exists. Otherwise, it creates a database with the Id
func (c *DocClient) CreateDatabaseIfNotExist(dbID string) (*Database, error) {
	db, err := c.GetDatabase(dbID)
	if err != nil {
		return nil, err
	}

	if db != nil {
		return db, nil
	}

	db, err = c.CreateDatabase(dbID)
	if err != nil {
		return nil, err
	}

	return db, nil
}
