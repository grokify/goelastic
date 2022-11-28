package main

import (
	"fmt"
	"net/http"

	"github.com/grokify/gohttp/httpsimple"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"

	elastirad "github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/docs/reference"
	"github.com/grokify/elastirad-go/models/es5"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html#mappings

// main is a simple request that shows creating an index in action.
func main() {
	esClient, err := elastirad.NewSimpleClient("", "", "", true)
	logutil.FatalErr(err)

	body := es5.CreateIndexBody{
		Mappings: map[string]es5.Mapping{
			"tweet": {
				All: es5.All{Enabled: true},
				Properties: map[string]es5.Property{
					"message": {Type: "text"},
					"user": {
						Type: "nested",
						Properties: map[string]es5.Property{
							"username": {Type: "keyword"},
						}}}}}}

	fmtutil.MustPrintJSON(body)

	resp, err := esClient.Do(httpsimple.SimpleRequest{
		Method: http.MethodPut,
		URL:    "twitter",
		IsJSON: true,
		Body:   body})
	reference.ProcResponse(resp, err)

	fmt.Println("DONE")
}
