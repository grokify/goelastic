package main

import (
	"fmt"
	"net/url"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	v5 "github.com/grokify/elastirad-go/models/v5"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html

// main shows bool query in action
func main() {
	esClient := elastirad.NewClient(url.URL{})

	body := v5.QueryBody{
		Query: v5.Query{
			Bool: &v5.BoolQuery{
				Should: []v5.Filter{
					{Match: map[string]string{"hash_tags": "wow"}},
					{Match: map[string]string{"hash_tags": "elasticsearch"}}},
				MinimumShouldMatch: 1}}}

	esReq := models.Request{
		Method: "POST",
		Path:   []interface{}{"twitter/tweet", elastirad.SearchSlug},
		Body:   body}

	fmtutil.PrintJSON(body)

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
