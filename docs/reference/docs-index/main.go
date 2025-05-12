package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"

	"github.com/grokify/goelastic"
	"github.com/grokify/goelastic/docs/reference"
	"github.com/grokify/goelastic/models"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html

// main is a simple request that shows the ES documented
// index document request. After running this code, verify
// by checking http://localhost:9200/twitter/_search
func main() {
	esClient, err := goelastic.NewSimpleClient("", "", "", true)
	logutil.FatalErr(err)

	id := "1"
	tweet := reference.Tweet{
		User:     reference.User{Username: "kimchy"},
		PostDate: time.Now().Format(time.RFC3339),
		Message:  "trying out Elasticsearch",
		HashTags: []string{"elasticsearch", "wow"}}

	ctx := context.Background()

	// Create Doc
	resp, err := esClient.Do(ctx, httpsimple.Request{
		Method:   http.MethodPost,
		URL:      strings.Join([]string{"twitter/tweet", id, goelastic.SlugCreate}, "/"),
		BodyType: httpsimple.BodyTypeJSON,
		Body:     tweet})
	reference.ProcResponse(resp, err)

	// Get/Check Doc
	resp, err = esClient.Do(ctx, httpsimple.Request{
		Method: http.MethodGet,
		URL:    strings.Join([]string{"twitter/tweet", id}, "/")})
	reference.ProcResponse(resp, err)

	// update Doc
	tweet.Message = "trying out Elasticsearch again"

	resp, err = esClient.Do(ctx, httpsimple.Request{
		Method:   http.MethodPost,
		URL:      strings.Join([]string{"twitter/tweet", id, goelastic.SlugUpdate}, "/"),
		BodyType: httpsimple.BodyTypeJSON,
		Body:     models.UpdateIndexDoc{Doc: tweet}})
	reference.ProcResponse(resp, err)

	// Get/Check Doc
	resp, err = esClient.Do(ctx, httpsimple.Request{
		Method: http.MethodGet,
		URL:    strings.Join([]string{"twitter/tweet", id}, "/")})
	reference.ProcResponse(resp, err)

	fmt.Println("DONE")
}
