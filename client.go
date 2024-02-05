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
	ErrClientNotSet     = errors.New("httpsimple.Client not set")
	ErrDocumentIDNotSet = errors.New("document id s be empty")
	ErrRequestNotSet    = errors.New("request cannot be empty")
	ErrTargetNetSet     = errors.New("target cannot be empty")
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
func (c *Client) DocumentRead(target, id string, v any) (*DocumentReadAPIResponse, *http.Response, error) {
	apiResp := &DocumentReadAPIResponse{}
	if err := c.validateClientAndTarget(target); err != nil {
		return nil, nil, err
	} else if strings.TrimSpace(id) == "" {
		return nil, nil, ErrDocumentIDNotSet
	} else if resp, err := c.SimpleClient.Do(httpsimple.Request{
		Method: http.MethodGet,
		URL:    urlutil.JoinAbsolute(target, SlugDoc, id),
	}); err != nil {
		return nil, resp, err
	} else if b, err := io.ReadAll(resp.Body); err != nil {
		return nil, resp, err
	} else if err = json.Unmarshal(b, apiResp); err != nil {
		return apiResp, resp, err
	} else if v == nil {
		return apiResp, resp, nil
	} else {
		return apiResp, resp, json.Unmarshal(apiResp.Source, v)
	}
}

// DocumentReadAPIResponse represents an Elasticsearch document response.
type DocumentReadAPIResponse struct {
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
	} else if len(username) > 0 || len(password) > 0 {
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
