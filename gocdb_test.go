package gocdb

import (
	"testing"
)

func TestAuthSigGeneration(t *testing.T) {
	v := "GET"
	rt := "dbs"
	ri := "dbs/ToDoList"
	d := "Thu, 27 Apr 2017 00:51:12 GMT"
	k := "dsZQi3KtZmCv1ljt3VNWNm7sQUF1y5rJfC6kv5JiwvW0EndXdDku/dkKBp8/ufDToSxLzR4y+O/0H/t4bQtVNw=="

	// expected, _ := url.QueryUnescape("type%3dmaster%26ver%3d1.0%26sig%3dc09PEVJrgp2uQRkr934kFbTqhByc7TVr3OHyqlu%2bc%2bc%3d")
	expected := "type%3Dmaster%26ver%3D1.0%26sig%3Dc09PEVJrgp2uQRkr934kFbTqhByc7TVr3OHyqlu%2Bc%2Bc%3D"

	s := generateAuthSig(v, rt, ri, d, k, "master", "1.0")
	// s, _ = url.QueryUnescape(s)
	if s != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, s)
	}
}
