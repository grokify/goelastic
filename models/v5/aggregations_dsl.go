package v5

import (
	"encoding/json"
)

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
	Size  int            `json:"size"`
	Query *Query         `json:"query,omitempty"`
	Aggs  map[string]Agg `json:"aggs,omitempty"`
}

// Agg represents one aggregation.
// See https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-filter-aggregation.html
// DateHistogram: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-datehistogram-aggregation.html#search-aggregations-bucket-datehistogram-aggregation
type Agg struct {
	Filter        *Filter        `json:"filter,omitempty"`
	Nested        *Nested        `json:"nested,omitempty"`
	Terms         *Term          `json:"terms,omitempty"`
	Aggs          map[string]Agg `json:"aggs,omitempty"`
	DateHistogram *DateHistogram `json:"date_histogram,omitempty"`
}

// DateHistogram represents the Elasticsearch date_histogram aggregation.
type DateHistogram struct {
	Field    string `json:"field,omitempty"`
	Interval string `json:"interval,omitempty"`
}

// Nested indicates whether the aggregation key is
// a nested key or not. It is mandatory to use a
// nested key (e.g. a key using dot notation).
type Nested struct {
	Path      string `json:"path,omitempty"`
	ScoreMode string `json:"score_mode,omitempty"`
	Query     *Query `json:"query,omitempty"`
}

// Term is the field being processed.
type Term struct {
	Field string `json:"field,omitempty"`
	Null  bool   `json:"null,omitempty"`
}

type AggResponseSimple struct {
	Aggregations map[string]Aggregation `json:"aggregations,omitempty"`
}

// AggResponseWIP is working response object used to
// format an aggregations response using Elastirad's
// AggResponseRad struct.
type AggResponseNestedWIP struct {
	Aggregations map[string]map[string]interface{} `json:"aggregations,omitempty"`
}

// AggregationResRad is a Go optimized response struct.
type AggregationResRad struct {
	AggregationName string      `json:"aggregation_name,omitempty"`
	Type            string      `json:"type,omitempty"`
	DocCount        int         `json:"doc_count,omitempty"`
	AggregationData Aggregation `json:"aggregation_data,omitempty"`
}

// Aggregation is a Elasticsearch aggregation response struct.
type Aggregation struct {
	DocCountErrorUpperBound int      `json:"doc_count_error_upper_bound,omitempty"`
	SumOtherDocCount        int      `json:"sum_other_doc_count,omitempty"`
	Buckets                 []Bucket `json:"buckets,omitempty"`
}

// Bucket is an Elasticsearch aggregation response bucket struct.
type Bucket struct {
	Key         interface{} `json:"key,omitempty"`
	KeyAsString string      `json:"key_as_string,omitempty"`
	DocCount    int         `json:"doc_count,omitempty"`
}

// AggregationResRadArrayFromBodyBytesNested returns an array of formatted
// Aggregations using AggregationResRad structs given a HTTP response byte
// array.
// Nested
func AggregationResRadArrayFromBodyBytesNested(bytes []byte) ([]AggregationResRad, error) {
	esRes := AggResponseNestedWIP{}
	esAggs := []AggregationResRad{}
	err := json.Unmarshal(bytes, &esRes)
	if err != nil {
		return esAggs, err
	}
	for k1, v1 := range esRes.Aggregations {
		esAgg := AggregationResRad{Type: k1}
		for k2, v2 := range v1 {
			if k2 == "doc_count" {
				esAgg.DocCount = int(v2.(float64))
			} else {
				esAgg.AggregationName = k2
				aggregationJSONBytes, err := json.Marshal(v2)
				if err != nil {
					return esAggs, err
				}
				agg := Aggregation{}
				err = json.Unmarshal(aggregationJSONBytes, &agg)
				if err != nil {
					return esAggs, err
				}
				esAgg.AggregationData = agg
			}
		}
		esAggs = append(esAggs, esAgg)
	}
	return esAggs, nil
}

// AggregationResRadArrayFromBodyBytes is used to parse
// simple, non-nested aggregations.
func AggregationResRadArrayFromBodyBytes(bytes []byte) ([]AggregationResRad, error) {
	esRes := AggResponseSimple{}
	esAggs := []AggregationResRad{}
	err := json.Unmarshal(bytes, &esRes)
	if err != nil {
		return esAggs, err
	}
	for k1, srcAgg := range esRes.Aggregations {
		esAgg := AggregationResRad{
			AggregationName: k1,
			DocCount:        0,
			AggregationData: srcAgg,
		}
		for _, bucket := range srcAgg.Buckets {
			esAgg.DocCount += bucket.DocCount
		}
		esAggs = append(esAggs, esAgg)
	}
	return esAggs, nil
}
