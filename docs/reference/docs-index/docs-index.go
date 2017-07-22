package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/docs/reference"
	"github.com/grokify/elastirad-go/models"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html

func createDoc(esClient elastirad.Client, id string, doc interface{}) {
	esReq := models.Request{
		Method: "POST",
		Path:   []interface{}{"twitter/tweet", id, elastirad.CreateSlug},
		Body:   doc}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Printf("C_ERR: %v\n", err)
	} else {
		fmt.Printf("C_RES_BODY: %v\n", string(res.Body()))
		fmt.Printf("C_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}

func getDoc(esClient elastirad.Client, id string) {
	esReq := models.Request{
		Method: "GET",
		Path:   []interface{}{"twitter/tweet", id}}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Printf("R_ERR: %v\n", err)
	} else {
		fmt.Printf("R_RES_BODY: %v\n", string(res.Body()))
		fmt.Printf("R_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}

func updateDoc(esClient elastirad.Client, id string, doc interface{}) {
	esReq := models.Request{
		Method: "POST",
		Path:   []interface{}{"twitter/tweet", id, elastirad.UpdateSlug},
		Body:   doc}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
	} else {
		fmt.Printf("U_RES_BODY: %v\n", string(res.Body()))
		fmt.Printf("U_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}

// main is a simple request that shows the ES documented
// index document request. After running this code, verify
// by checking http://localhost:9200/twitter/_search
func main() {
	esClient := elastirad.NewClient(url.URL{})

	id := "1"
	tweet := ref.Tweet{
		User:     ref.User{Username: "kimchy"},
		PostDate: time.Now().Format(time.RFC3339),
		Message:  "trying out Elasticsearch",
		HashTags: []string{"elasticsearch", "wow"}}

	createDoc(esClient, id, tweet)
	getDoc(esClient, id)
	tweet.Message = "trying out Elasticsearch again"
	updateDoc(esClient, id, models.UpdateIndexDoc{Doc: tweet})
	getDoc(esClient, id)
}
