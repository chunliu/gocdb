package gocdb

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DocClient represents the client to the Cosmos DB
type DocClient struct {
	Endpoint *url.URL
	AuthKey  string
}

func (c *DocClient) GetDatabase(dbId string) (Database, error) {
	ri := fmt.Sprintf("dbs/%s", dbId)
	url := fmt.Sprintf("%s/dbs/%s", c.Endpoint, dbId)
	utcDate := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	authSig := generateAuthSig("Get", "dbs", ri, utcDate, c.AuthKey, "master", "1.0")

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Add("x-ms-date", utcDate)
	req.Header.Set("authorization", authSig)
	req.Header.Add("x-ms-version", "2017-02-22")
	req.Header.Set("User-Agent", "gocdb/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Database{}, fmt.Errorf("Failed to get database. Error: %v", err)
	}

	return Database{CollectionsLink: resp.Status}, nil
}

// func (c *DocClient) CreateDatabaseIfNotExist(dbId string) Database {
// 	return nil, nil
// }
