package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"

	esj "github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
)

type Tweet struct {
	User     string `json:"user,omitempty"`
	PostDate string `json:"post_date,omitempty"`
	Message  string `json:"message,omitempty"`
}

// main is a simple request that shows the ES documented
// index document request. After running this code, verify
// by checking http://localhost:9200/twitter/_search
func main() {
	esClient := esj.NewClient(url.URL{})

	tweet := Tweet{
		User:     "kimchy",
		PostDate: time.Now().Format(time.RFC3339),
		Message:  "trying out Elasticsearch"}

	esReq := models.Request{
		Method: "POST",
		Path:   []interface{}{"twitter/tweet", 1, "_create"},
		Body:   tweet}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	} else {
		fmt.Printf("RES: %v\n", res.StatusCode())
	}

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
