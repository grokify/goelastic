package es5

// https://www.elastic.co/guide/en/elasticsearch/reference/current/_executing_aggregations.html

/*
{
  "mappings": {
    "data": {
      "properties": {
        "state" : {
          "type": "string",
          "fields": {
            "raw" : {
              "type": "string",
              "index": "not_analyzed"
            }
          }
        }
      }
    }
  }
}

PUT my_index
{
  "mappings": {
    "my_type": {
      "properties": {
        "tags": {
          "type":  "keyword"
        }
      }
    }
  }
}

*/

type CreateIndexBody struct {
	Index    *Index             `json:"index,omitempty"`
	Settings *Settings          `json:"settings,omitempty"`
	Mappings map[string]Mapping `json:"mappings,omitempty"`
}

type Index struct {
	NumberOfShards   uint32 `json:"number_of_shards,omitempty"`
	NumberOfReplicas uint32 `json:"number_of_replicas,omitempty"`
}

type Settings struct {
	NumberOfShards int `json:"number_of_shards,omitempty"`
}

type Mapping struct {
	All        All                 `json:"_all,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
}

type All struct {
	Enabled bool `json:"enabled,omitempty"`
}

type Property struct {
	Type       string              `json:"type,omitempty"`
	Index      string              `json:"index,omitempty"`
	Format     string              `json:"format,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
	Fields     map[string]Property `json:"fields,omitempty"` // key can be "raw"
}
