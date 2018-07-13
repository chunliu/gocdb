package gocdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func (c *DocumentClient) CreateDocumentCollection(dbID string, coll *Collection, options RequestOptions) (*Collection, error) {
	result := make(chan *cdbError)

	go func() {
		verb := "POST"
		ri := fmt.Sprintf("dbs/%s", dbID)
		url := fmt.Sprintf("%s/dbs/%s/colls", c.Endpoint, dbID)
		cv, _ := json.Marshal(coll)

		req, err := createRequest(verb, url, "colls", ri, c.AuthKey, bytes.NewBuffer(cv))
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("CreateDocumentCollection: Failed to create request. Error: %v", err),
			}
			return
		}

		if options.OfferThroughput > 0 {
			req.Header.Add("x-ms-offer-throughput", strconv.Itoa(options.OfferThroughput))
		}
		if len(options.OfferType) > 0 {
			req.Header.Add("x-ms-offer-type", options.OfferType)
		}

		dbe := sendCdbRequest(req, coll)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("CreateDocumentCollection: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	if r.data == nil {
		return nil, r.err
	}

	return r.data.(*Collection), nil
}

func (c *DocumentClient) GetDocumentCollection(dbID, collID string) (*Collection, error) {
	result := make(chan *cdbError)

	go func() {
		verb := "GET"
		ri := fmt.Sprintf("dbs/%s/colls/%s", dbID, collID)
		url := fmt.Sprintf("%s/%s", c.Endpoint, ri)

		req, err := createRequest(verb, url, "colls", ri, c.AuthKey, nil)
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("GetDocumentCollection: Failed to create request. Error: %v", err),
			}
			return
		}

		coll := &Collection{}
		dbe := sendCdbRequest(req, coll)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("GetDocumentCollection: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	if r.data == nil {
		return nil, r.err
	}

	return r.data.(*Collection), nil
}

func (c *DocumentClient) CreateDocumentCollectionIfNotExist(dbID string, coll *Collection, options RequestOptions) (*Collection, error) {
	collection, err := c.GetDocumentCollection(dbID, coll.Id)

	if err != nil {
		return nil, err
	}
	if collection != nil {
		return collection, nil
	}

	return c.CreateDocumentCollection(dbID, coll, options)
}

func (c *DocumentClient) DeleteDocumentCollection(dbID, collID string) error {
	result := make(chan *cdbError)

	go func() {
		verb := "DELETE"
		ri := fmt.Sprintf("dbs/%s/colls/%s", dbID, collID)
		url := fmt.Sprintf("%s/%s", c.Endpoint, ri)

		req, err := createRequest(verb, url, "colls", ri, c.AuthKey, nil)
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("DeleteDocumentCollection: Failed to create request. Error: %v", err),
			}
			return
		}

		coll := &Collection{}
		dbe := sendCdbRequest(req, coll)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("DeleteDocumentCollection: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result

	return r.err
}

func (c *DocumentClient) ReplaceDocumentCollection(dbID string, coll *Collection) (*Collection, error) {
	result := make(chan *cdbError)

	go func() {
		verb := "PUT"
		ri := fmt.Sprintf("dbs/%s/colls/%s", dbID, coll.Id)
		url := fmt.Sprintf("%s/%s", c.Endpoint, ri)
		cv, _ := json.Marshal(coll)

		req, err := createRequest(verb, url, "colls", ri, c.AuthKey, bytes.NewBuffer(cv))
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("ReplaceDocumentCollection: Failed to create request. Error: %v", err),
			}
			return
		}

		dbe := sendCdbRequest(req, coll)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("ReplaceDocumentCollection: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	if r.data == nil {
		return nil, r.err
	}

	return r.data.(*Collection), nil
}
