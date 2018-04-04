package gocdb

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Get time.Now in RFC7231 (HTTP-date) format
func utcNow() string {
	return time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT") // RFC 7231: HTTP-date
}

// Generate authorization signature for Cosmos DB
func generateAuthSig(verb, resourceType, resourceID, utcDate, key, keyType, version string) string {
	decodedKey, _ := base64.StdEncoding.DecodeString(key)
	hmacSha256 := hmac.New(sha256.New, decodedKey)
	//utcDate := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT") // RFC 7231: HTTP-date

	payload := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		strings.ToLower(verb),
		strings.ToLower(resourceType),
		resourceID,
		strings.ToLower(utcDate),
		"",
	)

	hmacSha256.Write([]byte(payload))
	hashPayload := hmacSha256.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(hashPayload)

	v := fmt.Sprintf("type=%s&ver=%s&sig=%s", keyType, version, signature)

	// Signature need to be in URL encoding
	return url.QueryEscape(v)
}

func createRequest(verb, url, utcDate, authSig string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("x-ms-date", utcDate)
	req.Header.Set("authorization", authSig)
	req.Header.Add("x-ms-version", "2017-02-22")
	req.Header.Set("User-Agent", "gocdb/1.0")

	return req, nil
}
