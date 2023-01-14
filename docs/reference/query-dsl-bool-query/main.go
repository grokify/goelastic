package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"

	elastirad "github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/docs/reference"
	"github.com/grokify/elastirad-go/models/es5"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html

// main shows bool query in action
func main() {
	esClient, err := elastirad.NewSimpleClient("", "", "", true)
	logutil.FatalErr(err)

	body := es5.QueryBody{
		Query: es5.Query{
			Bool: &es5.BoolQuery{
				Should: []es5.Filter{
					{Match: map[string]string{"hash_tags": "wow"}},
					{Match: map[string]string{"hash_tags": "elasticsearch"}}},
				MinimumShouldMatch: 1}}}

	fmtutil.MustPrintJSON(body)

	resp, err := esClient.Do(httpsimple.SimpleRequest{
		Method:   http.MethodPost,
		URL:      strings.Join([]string{"twitter/tweet", elastirad.SearchSlug}, "/"),
		BodyType: httpsimple.BodyTypeJSON,
		Body:     body})
	reference.ProcResponse(resp, err)

	fmt.Println("DONE")
}
