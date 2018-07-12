package gocdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// DocClient represents the client to the Cosmos DB
type DocClient struct {
	Endpoint *url.URL
	AuthKey  string
}

type cdbError struct {
	data interface{}
	err  error
}

// GetDatabase restrieves the database info based on the database Id
func (c *DocClient) GetDatabase(dbID string) (*Database, error) {
	result := make(chan *cdbError)

	go func() {
		verb := "GET"
		ri := fmt.Sprintf("dbs/%s", dbID)
		url := fmt.Sprintf("%s/dbs/%s", c.Endpoint, dbID)

		req, err := createRequest(verb, url, "dbs", ri, c.AuthKey, nil)
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("GetDatabase: Failed to create request. Error: %v", err),
			}
			return
		}

		db := &Database{}
		dbe := sendCdbRequest(req, db)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("GetDatabase: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	if r.data == nil {
		return nil, r.err
	}

	return r.data.(*Database), r.err
}

// CreateDatabase create a database with the Id
func (c *DocClient) CreateDatabase(dbID string) (*Database, error) {
	result := make(chan *cdbError)

	go func() {
		verb := "POST"
		url := fmt.Sprintf("%s/dbs", c.Endpoint)

		db := &Database{}
		db.Id = dbID

		jv, _ := json.Marshal(db)

		req, err := createRequest(verb, url, "dbs", "", c.AuthKey, bytes.NewBuffer(jv))
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("CreateDatabase: Failed to create request. Error: %v", err),
			}
			return
		}

		dbe := sendCdbRequest(req, db)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("CreateDatabase: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	if r.data == nil {
		return nil, r.err
	}

	return r.data.(*Database), r.err
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

	// Not found
	return c.CreateDatabase(dbID)
}

// ListDatabases retrieve all databases from the remote
func (c *DocClient) ListDatabases() (*Databases, error) {
	result := make(chan *cdbError)
	go func() {
		verb := "GET"
		url := fmt.Sprintf("%s/dbs", c.Endpoint)

		req, err := createRequest(verb, url, "dbs", "", c.AuthKey, nil)
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("ListDatabases: Failed to create request. Error: %v", err),
			}
			return
		}

		dbs := &Databases{}
		dbe := sendCdbRequest(req, dbs)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("ListDatabases: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	if r.data == nil {
		return nil, r.err
	}

	return r.data.(*Databases), r.err
}

// DeleteDatabase delete the database from remote
func (c *DocClient) DeleteDatabase(dbID string) error {
	result := make(chan *cdbError)

	go func() {
		verb := "DELETE"
		ri := fmt.Sprintf("dbs/%s", dbID)
		url := fmt.Sprintf("%s/dbs/%s", c.Endpoint, dbID)

		req, err := createRequest(verb, url, "dbs", ri, c.AuthKey, nil)
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("DeleteDatabase: Failed to create request. Error: %v", err),
			}
			return
		}

		db := &Database{}
		dbe := sendCdbRequest(req, db)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("DeleteDatabase: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result

	return r.err
}
