package models

import (
	"net/url"
)

// Request is an elastirad.Request which contains generic
// information for an Elasticsearch API request.
type Request struct {
	Method      string
	Path        []interface{}
	Query       url.Values
	ContentType string
	Body        interface{}
}

// UpdateIndexDoc represents the body for a doc partial update.
type UpdateIndexDoc struct {
	Doc         interface{} `json:"doc,omitempty"`
	DocAsUpsert bool        `json:"doc_as_upsert,omitempty"`
}

// CreateIndex represents to create index API request body.
type CreateIndex struct {
	Settings Settings           `json:"settings,omitempty"`
	Mappings map[string]Mapping `json:"mappings,omitempty"`
	Aliases  map[string]Alias   `json:"aliases,omitempty"`
}

// Settings is the struct for the create index API request body parameter.
type Settings struct {
	Index                    Index
	NumberOfShards           int    `json:"number_of_shares,omitempty"`
	NumberOfReplicas         int    `json:"number_of_replicas,omitempty"`
	WriteWaitForActiveShards string `json:"index.write.wait_for_active_shards,omitempty"`
}

// Index is the metadata for the search index.
type Index struct {
	NumberOfShards   int `json:"number_of_shares,omitempty"`
	NumberOfReplicas int `json:"number_of_replicas,omitempty"`
}

// Mapping is the struct for the create index API request body parameter.
type Mapping struct {
}

// Alias is the struct for the create index API request body parameter.
type Alias struct {
	Routing string `json:"routing,omitempty"`
}
