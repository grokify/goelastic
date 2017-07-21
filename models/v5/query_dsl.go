package v5

type QueryRequest struct {
	Query Query `json:"query,omitempty"`
}

type Query struct {
	Bool BoolQuery `json:"bool,omitempty"`
}

type BoolQuery struct {
	Must               []Filter `json:"must,omitempty"`
	Should             []Filter `json:"should,omitempty"`
	MinimumShouldMatch int      `json:"minimum_should_match,omitempty"`
	Boost              float64  `json:"boost,omitempty"`
}

type Filtered struct {
	Or FilteredOr `json:"or,omitempty"`
}

type FilteredOr struct {
	Filters []Filter `json:"filters,omitempty"`
	Cache   bool     `json:"_cache,qromitempty"`
}

type Filter struct {
	Match map[string]string `json:"match,omitempty"`
	Term  map[string]string `json:"term,omitempty"`
}
