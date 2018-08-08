package gocdb

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *DocumentClient) CreateDocument(dbID, collID string, doc interface{}, options RequestOptions) (interface{}, error) {
	result := make(chan *cdbError)

	go func() {
		verb := "POST"
		ri := fmt.Sprintf("dbs/%s/colls/%s", dbID, collID)
		url := fmt.Sprintf("%s/%s/docs", c.Endpoint, ri)
		dv, _ := json.Marshal(doc)

		req, err := createRequest(verb, url, "docs", ri, c.AuthKey, bytes.NewBuffer(dv))
		if err != nil {
			result <- &cdbError{
				data: nil,
				err:  fmt.Errorf("CreateDocument: Failed to create request. Error: %v", err),
			}
			return
		}

		if options.IsUpsert {
			req.Header.Add("x-ms-documentdb-is-upsert", "true")
		}
		if len(options.IndexingDirective) > 0 {
			req.Header.Add("x-ms-indexing-directive", options.IndexingDirective)
		}

		dbe := sendCdbRequest(req, doc)
		if dbe.err != nil {
			dbe.err = fmt.Errorf("CreateDocument: %v", dbe.err)
		}
		result <- dbe
	}()

	r := <-result
	if r.data == nil {
		return nil, r.err
	}

	return r.data, nil
}
