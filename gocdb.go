package gocdb

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

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
