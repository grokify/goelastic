package v5

// https://www.elastic.co/guide/en/elasticsearch/reference/current/_executing_aggregations.html

/*
GET /bank/_search
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword"
      }
    }
  }
}

{
    "aggs" : {
        "genres" : {
            "terms" : { "field" : "genre" }
        }
    }
}
*/

// AggsBody represents the body of an aggregation
// API request.
type AggsBody struct {
	Size int            `json:"size"`
	Aggs map[string]Agg `json:"aggs,omitempty"`
}

// Agg represents one aggregation.
type Agg struct {
	Nested *Nested        `json:"nested,omitempty"`
	Terms  *Term          `json:"terms,omitempty"`
	Aggs   map[string]Agg `json:"aggs,omitempty"`
}

// Nested indicates whether the aggregation key is
// a nested key or not. It is mandatory to use a
// nested key (e.g. a key using dot notation).
type Nested struct {
	Path string `json:"path,omitempty"`
}

// Term is the field being processed.
type Term struct {
	Field string `json:"field,omitempty"`
	Null  bool   `json:"null,omitempty"`
}
