package gocdb

type Resource struct {
	Id         string `json:"id"`
	ResourceId string `json:"_rid"`
	SelfLink   string `json:"_self"`
	ETag       string `json:"_etag"`
	Timestamp  string `json:"_ts"`
}

type Database struct {
	Resource
	CollectionsLink string `json:"_coll"`
	UsersLink       string `json:"_users"`
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
