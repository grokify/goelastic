# Elastirad: A Go Client Wrapper for Elasticsearch

[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

Simple client to query Elasticsearch API using HTTP API documentation. This is inspired by the [Ruby Elastirad gem](https://github.com/grokify/elastirad-ruby).

## Usage

### Index API

* [Index Document](examples/index_document) ([ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html))

### Bool Query

Here is a Bool Query example from [Elasticsearch 5.5: Bool Query](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html)

```golang

import (
	"github.com/grokify/elastirad-go/models"
	"github.com/grokify/elastirad-go/models/v5"
)

func main() {
	qry := v5.QueryBody{
		Query: v5.Query{
			Bool: v5.BoolQuery{
				Should: []v5.Filter{
					{Match: map[string]string{"tag": "wow"}},
					{Match: map[string]string{"tag": "elasticsearch"}}},
				MinimumShouldMatch: 1}}}

	req := models.Request{
		Method: "POST",
		Path:   []interface{}{"twitter/tweet", elastirad.SearchSlug},
		Body:   qry}
}
```

Which results in the following query:

```json
{
  "query": {
    "bool" : {
      "should" : [
        { "term" : { "tag" : "wow" } },
        { "term" : { "tag" : "elasticsearch" } }
      ],
      "minimum_should_match" : 1
    }
  }
}
```

 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/elastirad-go
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/elastirad-go
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/elastirad-go
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/elastirad-go/blob/master/LICENSE.md