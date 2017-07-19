package models

import (
	"net/url"
)

type Request struct {
	Method      string
	Path        []interface{}
	Query       url.Values
	ContentType string
	Body        interface{}
}

type CreateIndex struct {
	Settings Settings           `json:"settings,omitempty"`
	Mappings map[string]Mapping `json:"mappings,omitempty"`
	Aliases  map[string]Alias   `json:"aliases,omitempty"`
}

type Settings struct {
	Index                    Index
	NumberOfShards           int    `json:"number_of_shares,omitempty"`
	NumberOfReplicas         int    `json:"number_of_replicas,omitempty"`
	WriteWaitForActiveShards string `json:"index.write.wait_for_active_shards,omitempty"`
}
type Index struct {
	NumberOfShards   int `json:"number_of_shares,omitempty"`
	NumberOfReplicas int `json:"number_of_replicas,omitempty"`
}

type Mapping struct {
}

type Alias struct {
	Routing string `json:"routing,omitempty"`
}
