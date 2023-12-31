package es5

/*
// Request is an geneic request which contains generic
// information for an Elasticsearch API request.
type Request struct {
	Method      string
	Path        []interface{}
	Query       url.Values
	ContentType string
	Body        interface{}
}
*/

type ResponseBody struct {
	Hits Hits `json:"hits,omitempty"`
}

type Hits struct {
	Total int64 `json:"total,omitempty"`
}
