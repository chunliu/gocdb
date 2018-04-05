package gocdb

type Resource struct {
	Id         string `json:"id"`
	ResourceId string `json:"_rid,omitempty"`
	SelfLink   string `json:"_self,omitempty"`
	ETag       string `json:"_etag,omitempty"`
	Timestamp  int    `json:"_ts,omitempty"`
}

type Database struct {
	Resource
	CollectionsLink string `json:"_coll,omitempty"`
	UsersLink       string `json:"_users,omitempty"`
}

type Databases struct {
	ResourceId string `json:"_rid,omitempty"`
	Databases  []Database
	Count      int `json:"_count,omitempty"`
}

type PartitionKey struct {
	Paths []string `json:"paths"`
}

type Collection struct {
	Resource
	DefaultTimeToLive    int    `json:"defaultTtl"`
	DocumentsLink        string `json:"_docs"`
	PartitionKey         `json:"partitionKey"`
	StoredProceduresLink string `json:"_sprocs"`
}
