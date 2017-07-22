package v5

import (
	"net/url"
)

// Request is an elastirad.Request which contains generic
// information for an Elasticsearch API request.
type Request struct {
	Method      string
	Path        []interface{}
	Query       url.Values
	ContentType string
	Body        interface{}
}
