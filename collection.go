package gocdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func (c *DocClient) CreateDocumentCollection(dbID string, coll *Collection, options RequestOptions) (*Collection, error) {
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
