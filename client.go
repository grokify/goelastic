package elastirad

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/grokify/goauth/authutil"
	"github.com/grokify/mogo/net/http/httpsimple"
)

const (
	// ElasticsearchAPIDefaultScheme is the HTTP scheme for the default server.
	ElasticsearchAPIDefaultScheme string = "https"
	// ElasticsearchAPIDefaultHost is the HTTP host for the default server.
	ElasticsearchAPIDefaultHost string = "127.0.0.1:9200"
	// ElasticsearchAPIDefaultHost is the HTTP host for the default server.
	ElasticsearchAPIDefaultURL string = "https://127.0.0.1:9200"
	// CreateSlug is the URL path part for creates.
	CreateSlug string = "_create"
	// UpdateSlug is the URL path part for updates.
	UpdateSlug string = "_update"
	// SearchSlug is the URL path part for search.
	SearchSlug string = "_search"
)

func NewSimpleClient(serverURL, username, password string, allowInsecure bool) (httpsimple.SimpleClient, error) {
	if len(strings.TrimSpace(serverURL)) == 0 {
		serverURL = ElasticsearchAPIDefaultURL
	}
	if len(username) > 0 || len(password) > 0 {
		hclient, err := authutil.NewClientBasicAuth(username, password, allowInsecure)
		return httpsimple.SimpleClient{
			BaseURL:    serverURL,
			HTTPClient: hclient}, err
	}
	return httpsimple.SimpleClient{
		BaseURL:    serverURL,
		HTTPClient: authutil.NewClientHeaderQuery(http.Header{}, url.Values{}, allowInsecure)}, nil
}

/*

// Client is a API client for Elasticsearch.
type Client struct {
	BaseURL        url.URL
	Client         *http.Client
	FastHTTPClient *fasthttp.Client
}

// NewClient returns a Client struct given a Elasticsearch
// server URL.
func NewClient(baseURL url.URL) Client {
	c := Client{
		BaseURL:        baseURL,
		FastHTTPClient: &fasthttp.Client{}}
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

func (c *Client) BuildRequest(esReq models.Request) (*http.Request, error) {
	esURL := c.BuildURL(esReq)
	var body *bytes.Reader
	if esReq.Body != nil {
		data, err := json.Marshal(esReq.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(data)
	}
	return http.NewRequest(esReq.Method, esURL.String(), body)
}

func (c *Client) SendRequest(esReq models.Request) (*http.Response, error) {
	req, err := c.BuildRequest(esReq)
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(esReq.ContentType)) > 0 {
		req.Header.Add(httputilmore.HeaderContentType, esReq.ContentType)
	} else {
		req.Header.Add(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)
	}

	client := c.Client
	if client == nil {
		client = &http.Client{}
	}
	return client.Do(req)
}

var ErrEmptyRequest = errors.New("request cannot be empty")

func (c *Client) SendSimpleRequest(sreq *httpsimple.SimpleRequest) (*http.Response, error) {
	if sreq == nil {
		return nil, ErrEmptyRequest
	}
	if !urlutil.URIHasScheme(sreq.URL) {
		u := url.URL{
			Scheme: c.BaseURL.Scheme,
			Host:   c.BaseURL.Host,
			Path:   sreq.URL}
		sreq.URL = u.String()
	}
	return httpsimple.Do(*sreq)
}

// BuildFastRequest builds a valyala/fasthttp HTTP request struct.
func (c *Client) BuildFastRequest(esReq models.Request) (*fasthttp.Request, error) {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod(esReq.Method)
	esURL := c.BuildURL(esReq)
	req.Header.SetRequestURI(esURL.String())

	if len(strings.TrimSpace(esReq.ContentType)) > 0 {
		req.Header.Set(httputilmore.HeaderContentType, esReq.ContentType)
	} else {
		req.Header.Set(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)
	}

	if esReq.Body != nil {
		bytes, err := json.Marshal(esReq.Body)
		if err != nil {
			return req, err
		}
		req.SetBody(bytes)
	}

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

*/
