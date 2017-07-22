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

type AggsBody struct {
	Size int            `json:"size"`
	Aggs map[string]Agg `json:"aggs,omitempty"`
}

type Agg struct {
	Nested *Nested        `json:"nested,omitempty"`
	Terms  *Term          `json:"terms,omitempty"`
	Aggs   map[string]Agg `json:"aggs,omitempty"`
}

type Nested struct {
	Path string `json:"path,omitempty"`
}

type Term struct {
	Field string `json:"field,omitempty"`
	Null  bool   `json:"null,omitempty"`
}
