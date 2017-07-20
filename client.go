package elastirad

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/grokify/gotilla/net/httputil"
	"github.com/grokify/gotilla/strings/stringsutil"
	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go/models"
)

const (
	// ElasticsearchAPIDefaultScheme is the HTTP scheme for the default server.
	ElasticsearchAPIDefaultScheme string = "http"
	// ElasticsearchAPIDefaultHost is the HTTP host for the default server.
	ElasticsearchAPIDefaultHost string = "127.0.0.1:9200"
	// CreateSlug is the URL path part for creates.
	CreateSlug string = "_create"
	// UpdateSlug is the URL path part for updates.
	UpdateSlug string = "_update"
)

// Client is a API client for Elasticsearch.
type Client struct {
	BaseURL        url.URL
	FastHTTPClient fasthttp.Client
}

// NewClient returns a Client struct given a Elasticsearch
// server URL.
func NewClient(baseURL url.URL) Client {
	c := Client{
		BaseURL:        baseURL,
		FastHTTPClient: fasthttp.Client{}}
	c.SetDefaults()
	return c
}

// SetDefaults sets default values where not specified.
func (c *Client) SetDefaults() {
	if len(strings.TrimSpace(c.BaseURL.Scheme)) < 1 {
		c.BaseURL.Scheme = ElasticsearchAPIDefaultScheme
	}
	if len(strings.TrimSpace(c.BaseURL.Host)) < 1 {
		c.BaseURL.Host = ElasticsearchAPIDefaultHost
	}
}

// BuildFastRequest builds a valyala/fasthttp HTTP request struct.
func (c *Client) BuildFastRequest(esReq models.Request) (*fasthttp.Request, error) {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod(esReq.Method)
	esURL := c.BuildURL(esReq)
	req.Header.SetRequestURI(esURL.String())

	if len(strings.TrimSpace(esReq.ContentType)) > 0 {
		req.Header.Set(httputil.ContentTypeHeader, esReq.ContentType)
	} else {
		req.Header.Set(httputil.ContentTypeHeader, httputil.ContentTypeValueJSONUTF8)
	}

	bytes, err := json.Marshal(esReq.Body)
	if err != nil {
		return req, err
	}
	req.SetBody(bytes)

	return req, nil
}

// SendFastRequest executes a valyala/fasthttp HTTP request and returns
// the response, request and error structs.
func (c *Client) SendFastRequest(esReq models.Request) (*fasthttp.Response, *fasthttp.Request, error) {
	res := fasthttp.AcquireResponse()

	req, err := c.BuildFastRequest(esReq)
	if err != nil {
		return res, req, err
	}

	err = c.FastHTTPClient.Do(req, res)

	return res, req, err
}

// BuildURL merges the URL info in the request with the Elasticsearch
// server info configured in the client.
func (c *Client) BuildURL(esReq models.Request) url.URL {
	reqURL := url.URL{
		Scheme: c.BaseURL.Scheme,
		Host:   c.BaseURL.Host,
		Path:   stringsutil.JoinInterface(esReq.Path, "/", true, false, "")}
	return reqURL
}
