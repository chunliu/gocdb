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

type dbError struct {
	database *Database
	err      error
}

func sendDbRequest(req *http.Request) *dbError {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &dbError{
			database: nil,
			err:      err,
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return &dbError{nil, nil}
		}
		if resp.StatusCode != http.StatusCreated {
			return &dbError{
				nil,
				fmt.Errorf("HTTP Status: %v", resp.StatusCode),
			}
		}
	}

	db := &Database{}

	if err := json.NewDecoder(resp.Body).Decode(db); err != nil {
		return &dbError{
			database: nil,
			err:      err,
		}
	}

	return &dbError{
		database: db,
		err:      nil,
	}
}

// GetDatabase restrieves the database info based on the database Id
func (c *DocClient) GetDatabase(dbID string) (*Database, error) {
	result := make(chan *dbError)

	go func() {
		verb := "GET"
		ri := fmt.Sprintf("dbs/%s", dbID)
		url := fmt.Sprintf("%s/dbs/%s", c.Endpoint, dbID)
		utcDate := utcNow()
		authSig := generateAuthSig(verb, "dbs", ri, utcDate, c.AuthKey, "master", "1.0")

		req, err := createRequest(verb, url, utcDate, authSig, nil)
		if err != nil {
			result <- &dbError{
				database: nil,
				err:      fmt.Errorf("GetDatabase: Failed to create request. Error: %v", err),
			}
			return
		}

		dbe := sendDbRequest(req)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("GetDatabase: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	return r.database, r.err
}

// CreateDatabase create a database with the Id
func (c *DocClient) CreateDatabase(dbID string) (*Database, error) {
	result := make(chan *dbError)

	go func() {
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
			result <- &dbError{
				database: nil,
				err:      fmt.Errorf("CreateDatabase: Failed to create request. Error: %v", err),
			}
			return
		}

		dbe := sendDbRequest(req)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("CreateDatabase: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	return r.database, r.err
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
	db, err = c.CreateDatabase(dbID)
	if err != nil {
		return nil, err
	}

	return db, nil
}
