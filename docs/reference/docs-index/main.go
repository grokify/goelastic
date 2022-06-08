package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/grokify/gohttp/httpsimple"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/log/logutil"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/docs/reference"
	"github.com/grokify/elastirad-go/models"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html

// main is a simple request that shows the ES documented
// index document request. After running this code, verify
// by checking http://localhost:9200/twitter/_search
func main() {
	//esClient := elastirad.NewClient(url.URL{})
	esClient, err := elastirad.NewSimpleClient("", "", "", true)
	logutil.FatalErr(err)

	id := "1"
	tweet := reference.Tweet{
		User:     reference.User{Username: "kimchy"},
		PostDate: time.Now().Format(time.RFC3339),
		Message:  "trying out Elasticsearch",
		HashTags: []string{"elasticsearch", "wow"}}

	// Create Doc
	resp, err := esClient.Do(httpsimple.SimpleRequest{
		Method: http.MethodPost,
		URL:    strings.Join([]string{"twitter/tweet", id, elastirad.CreateSlug}, "/"),
		IsJSON: true,
		Body:   tweet})
	procResponse(resp, err)

	// Get/Check Doc
	resp, err = esClient.Do(httpsimple.SimpleRequest{
		Method: http.MethodGet,
		URL:    strings.Join([]string{"twitter/tweet", id}, "/")})
	procResponse(resp, err)

	// update Doc
	tweet.Message = "trying out Elasticsearch again"

	resp, err = esClient.Do(httpsimple.SimpleRequest{
		Method: http.MethodPost,
		URL:    strings.Join([]string{"twitter/tweet", id, elastirad.UpdateSlug}, "/"),
		IsJSON: true,
		Body:   models.UpdateIndexDoc{Doc: tweet}})
	procResponse(resp, err)

	// Get/Check Doc
	resp, err = esClient.Do(httpsimple.SimpleRequest{
		Method: http.MethodGet,
		URL:    strings.Join([]string{"twitter/tweet", id}, "/")})
	procResponse(resp, err)

	fmt.Println("DONE")
}

func procResponse(resp *http.Response, err error) {
	logutil.FatalErr(err)
	if resp == nil {
		return
	}
	body, err := jsonutil.PrettyPrintReader(resp.Body, "", "  ")
	logutil.FatalErr(err)
	fmt.Printf("C_RES_BODY: %v\n", string(body))
	fmt.Printf("C_RES_STATUS: %v\n", resp.StatusCode)
}
