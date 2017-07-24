package v5

// QueryBody represents a HTTP request body for the Elasticsearch
// query API.
type QueryBody struct {
	Query Query `json:"query,omitempty"`
}

// Query represents a Query API query object.
type Query struct {
	Bool   *BoolQuery            `json:"bool,omitempty"`
	Nested *Nested               `json:"nested,omitempty"`
	Match  map[string]MatchQuery `json:"match,omitempty"`
}

// MatchQuery represents a Query API match object.
type MatchQuery struct {
	Query          string `json:"query,omitempty"`
	Operator       string `json:"operator,omitempty"`
	ZeroTermsQuery string `json:"zero_terms_query,omitempty"`
}

// BoolQuery represents a Query API bool object.
type BoolQuery struct {
	Must               []Filter `json:"must,omitempty"`
	Should             []Filter `json:"should,omitempty"`
	MinimumShouldMatch int      `json:"minimum_should_match,omitempty"`
	Boost              float64  `json:"boost,omitempty"`
}

// Filtered represents a Query API filtered object.
type Filtered struct {
	Or FilteredOr `json:"or,omitempty"`
}

// FilteredOr represents a Query API filtered or object.
type FilteredOr struct {
	Filters []Filter `json:"filters,omitempty"`
	Cache   bool     `json:"_cache,qromitempty"`
}

// Filter represents a Query API filter object.
type Filter struct {
	Match map[string]string `json:"match,omitempty"`
	Term  map[string]string `json:"term,omitempty"`
}
