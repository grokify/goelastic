package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"

	"github.com/grokify/goelastic"
	"github.com/grokify/goelastic/docs/reference"
	"github.com/grokify/goelastic/models/es5"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html#mappings

// main is a simple request that shows creating an index in action.
func main() {
	esClient, err := goelastic.NewSimpleClient("", "", "", true)
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

	resp, err := esClient.Do(context.Background(), httpsimple.Request{
		Method:   http.MethodPut,
		URL:      "twitter",
		BodyType: httpsimple.BodyTypeJSON,
		Body:     body})
	reference.ProcResponse(resp, err)

	fmt.Println("DONE")
}
