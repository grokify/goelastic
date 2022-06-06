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
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html

// main shows bool query in action
func main() {
	esClient := elastirad.NewClient(url.URL{})

	body := es5.QueryBody{
		Query: es5.Query{
			Bool: &es5.BoolQuery{
				Should: []v5.Filter{
					{Match: map[string]string{"hash_tags": "wow"}},
					{Match: map[string]string{"hash_tags": "elasticsearch"}}},
				MinimumShouldMatch: 1}}}

	esReq := models.Request{
		Method: http.MethodPost,
		Path:   []interface{}{"twitter/tweet", elastirad.SearchSlug},
		Body:   body}

	fmtutil.MustPrintJSON(body)

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
