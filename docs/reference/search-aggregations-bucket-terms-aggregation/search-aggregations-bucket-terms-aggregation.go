package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	"github.com/grokify/elastirad-go/models/es5"
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
	esClient := elastirad.NewClient(url.URL{})

	body := es5.AggsBody{
		Size: 10000,
		Aggs: map[string]es5.Agg{
			"user": {
				Nested: &es5.Nested{Path: "user"},
				Aggs: map[string]es5.Agg{
					"TweetCountByUsername": {
						Terms: &es5.Term{Field: "user.username"}}}}}}

	fmtutil.MustPrintJSON(body)

	esReq := models.Request{
		Method: http.MethodPost,
		Path:   []interface{}{"twitter/tweet", elastirad.SearchSlug},
		Body:   body}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
	} else {
		fmt.Printf("U_RES_BODY: %v\n", string(res.Body()))
		fmt.Printf("U_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	fmt.Println("DONE")
}
