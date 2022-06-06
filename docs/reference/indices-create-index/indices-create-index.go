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
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html#mappings

// main is a simple request that shows creating an index
// in action.
func main() {
	esClient := elastirad.NewClient(url.URL{})

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

	esReq := models.Request{
		Method: http.MethodPut,
		Path:   []interface{}{"twitter"},
		Body:   body}

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
