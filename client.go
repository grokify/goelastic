package goelastic

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/grokify/goauth/authutil"
	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/net/urlutil"
)

var (
	ErrClientNotSet = errors.New("httpsimple.Client not set")
	ErrTargetNetSet = errors.New("target not set")
)

type Client struct {
	SimpleClient *httpsimple.Client
}

// IndexCreate creates an index. Documented at: https://www.elastic.co/guide/en/elasticsearch/reference/current/explicit-mapping.html .
func (c *Client) IndexCreate(target string, body any) (*http.Response, error) {
	if err := c.validateClientAndTarget(target); err != nil {
		return nil, err
	} else {
		return c.SimpleClient.Do(httpsimple.Request{
			Method:   http.MethodPut,
			URL:      target,
			Body:     body,
			BodyType: httpsimple.BodyTypeJSON,
		})
	}
}

// IndexPatch patches an index. Documented at: https://www.elastic.co/guide/en/elasticsearch/reference/current/explicit-mapping.html#add-field-mapping .
func (c *Client) IndexPatch(target string, body any) (*http.Response, error) {
	if err := c.validateClientAndTarget(target); err != nil {
		return nil, err
	} else {
		return c.SimpleClient.Do(httpsimple.Request{
			Method:   http.MethodPut,
			URL:      urlutil.JoinAbsolute(target, SlugMapping),
			Body:     body,
			BodyType: httpsimple.BodyTypeJSON,
		})
	}
}

// DocumentRead reads a document. If `v` is a pointer, the resulting `_source` property will be unmarshaled into it.
// The API is documented at: https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html and
// https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html#retrieving_document .
func (c *Client) DocumentRead(target, id string, v any) (*GetDocumentAPIResponse, *http.Response, error) {
	apiResp := &GetDocumentAPIResponse{}
	if err := c.validateClientAndTarget(target); err != nil {
		return nil, nil, err
	} else if resp, err := c.SimpleClient.Do(httpsimple.Request{
		Method: http.MethodGet,
		URL:    urlutil.JoinAbsolute(target, SlugDoc, id),
	}); err != nil {
		return nil, nil, err
	} else if b, err := io.ReadAll(resp.Body); err != nil {
		return nil, nil, err
	} else if err = json.Unmarshal(b, apiResp); err != nil {
		return nil, nil, err
	} else if v == nil {
		return apiResp, resp, nil
	} else {
		return apiResp, resp, json.Unmarshal(apiResp.Source, v)
	}
}

// GetDocumentAPIResponse represents an Elasticsearch document response.
type GetDocumentAPIResponse struct {
	Index  string          `json:"_index"`
	ID     string          `json:"_id"`
	Found  bool            `json:"found"`
	Source json.RawMessage `json:"_source"`
}

// DocumentRead reads a document. Documented at: https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html .
func (c *Client) DocumentReadSimple(target, id string) (*http.Response, error) {
	if err := c.validateClientAndTarget(target); err != nil {
		return nil, err
	} else {
		return c.SimpleClient.Do(httpsimple.Request{
			Method: http.MethodGet,
			URL:    urlutil.JoinAbsolute(target, SlugDoc, id),
		})
	}
}

// DocumentCreate crates a document with the document id `id`. If `id` is empty, a document id is created.`
func (c *Client) DocumentCreate(target, id string, body any) (*http.Response, error) {
	if err := c.validateClientAndTarget(target); err != nil {
		return nil, err
	} else {
		return c.SimpleClient.Do(httpsimple.Request{
			Method:   http.MethodPost,
			URL:      urlutil.JoinAbsolute(target, SlugCreate, id),
			Body:     body,
			BodyType: httpsimple.BodyTypeJSON,
		})
	}
}

func (c *Client) validateClientAndTarget(target string) error {
	if c.SimpleClient == nil {
		return ErrClientNotSet
	} else if strings.TrimSpace(target) == "" {
		return ErrTargetNetSet
	} else {
		return nil
	}
}

func NewSimpleClient(serverURL, username, password string, allowInsecure bool) (httpsimple.Client, error) {
	if len(strings.TrimSpace(serverURL)) == 0 {
		serverURL = DefaultServerURL
	}
	if len(username) > 0 || len(password) > 0 {
		hclient, err := authutil.NewClientBasicAuth(username, password, allowInsecure)
		return httpsimple.Client{
			BaseURL:    serverURL,
			HTTPClient: hclient}, err
	}
	return httpsimple.Client{
		BaseURL:    serverURL,
		HTTPClient: authutil.NewClientHeaderQuery(http.Header{}, url.Values{}, allowInsecure)}, nil
}

func NewConfigSimple(addrURL, username, password string, tlsInsecureSkipVerify bool) elasticsearch.Config {
	if strings.TrimSpace(addrURL) == "" {
		addrURL = DefaultServerURL
	}
	return elasticsearch.Config{
		Addresses: []string{
			addrURL,
		},
		Username: username,
		Password: password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: tlsInsecureSkipVerify, // #nosec G402 - used for local testing.
			},
		},
	}
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
