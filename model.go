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
	Kind  string   `json:"kind"`
}

type Index struct {
	Kind      string `json:"kind"`
	DataType  string `json:"dataType"`
	Precision int16  `json:"precision"`
}

type IncludedPath struct {
	Path    string `json:"path"`
	Indexes []Index
}

type ExcludedPath struct {
	Path string `json:"path"`
}

type IndexingPolicy struct {
	Automatic     bool
	IndexMode     string
	IncludedPaths []IncludedPath
	ExcludedPaths []ExcludedPath
}

type Collection struct {
	Resource
	PartitionKey         `json:"omitempty"`
	IndexingPolicy       `json:"omitempty"`
	DefaultTimeToLive    int    `json:"defaultTtl,omitempty"`
	DocumentsLink        string `json:"_docs,omitempty"`
	StoredProceduresLink string `json:"_sprocs,omitempty"`
	ConflictsLink        string `json:"_conflicts,omitempty"`
}

type DocumentBase struct {
	Resource
	Attachments string `json:"_attachments,omitempty"`
}

type RequestOptions struct {
	OfferThroughput   int
	OfferType         string
	IsUpsert          bool
	IndexingDirective string
}
