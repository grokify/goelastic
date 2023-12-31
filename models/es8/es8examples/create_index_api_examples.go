package es8examples

import (
	"net/http"

	"github.com/grokify/mogo/net/http/httpsimple"

	"github.com/grokify/goelastic/models/es8"
)

// CreateIndexExampleMappings provides a simple request for the example from: https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html#mappings
/*
Mappings
The create index API allows for providing a mapping definition:

PUT /test
{
  "settings": {
    "number_of_shards": 1
  },
  "mappings": {
    "properties": {
      "field1": { "type": "text" }
    }
  }
}
*/
func ExampleCreateIndexMappings() *httpsimple.Request {
	return &httpsimple.Request{
		Method:   http.MethodPut,
		URL:      "/test",
		BodyType: httpsimple.BodyTypeJSON,
		Body: es8.CreateIndexBody{
			Settings: &es8.Settings{
				NumberOfShards: 1,
			},
			Mappings: &es8.Mappings{
				Properties: map[string]es8.Property{
					"field1": {
						Type: "text",
					},
				},
			},
		},
	}
}
