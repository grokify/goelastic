package es8examples

import (
	"net/http"

	"github.com/grokify/elastirad-go/models/es8"
	"github.com/grokify/gohttp/httpsimple"
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
func ExampleCreateIndexMappings() *httpsimple.SimpleRequest {
	return &httpsimple.SimpleRequest{
		Method: http.MethodPut,
		URL:    "/test",
		IsJSON: true,
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
