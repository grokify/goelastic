package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"

	"github.com/grokify/goelastic"
	"github.com/grokify/goelastic/docs/reference"
	"github.com/grokify/goelastic/models/es5"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/5.5/_executing_aggregations.html
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-terms-aggregation.html#search-aggregations-bucket-terms-aggregation
// https://www.elastic.co/guide/en/elasticsearch/guide/current/aggregations-and-analysis.html#aggregations-and-analysis

// NESTED AGGREGATION
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-nested-aggregation.html
// main is a simple request that shows the ES documented
// request

// main shows nested aggregation in action
func main() {
	esClient, err := goelastic.NewSimpleClient("", "", "", true)
	logutil.FatalErr(err)

	body := es5.AggsBody{
		Size: 10000,
		Aggs: map[string]es5.Agg{
			"user": {
				Nested: &es5.Nested{Path: "user"},
				Aggs: map[string]es5.Agg{
					"TweetCountByUsername": {
						Terms: &es5.Term{Field: "user.username"}}}}}}

	fmtutil.MustPrintJSON(body)

	resp, err := esClient.Do(httpsimple.Request{
		Method:   http.MethodPost,
		URL:      strings.Join([]string{"twitter/tweet", goelastic.SlugSearch}, "/"),
		Body:     body,
		BodyType: httpsimple.BodyTypeJSON})
	reference.ProcResponse(resp, err)

	fmt.Println("DONE")
}
