# Elastirad: A Go Client Wrapper for Elasticsearch

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

Simple client to query Elasticsearch API using HTTP API documentation. This is inspired by the [Elastirad Ruby gem](https://github.com/grokify/elastirad-ruby).

## Usage

See the sample code in the [docs folder](docs).

So far the following example code has been created:

1. Create Index: [Go code](docs/reference/indices-create-index), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html)
1. Index Docs: [Go code](docs/reference/docs-index), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html)
1. Bool Query: [Go code](docs/reference/query-dsl-bool-query), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html)
1. Terms Aggregation: [Go code](docs/reference/search-aggregations-bucket-terms-aggregation), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-terms-aggregation.html#search-aggregations-bucket-terms-aggregation)

 [build-status-svg]: https://github.com/grokify/elastirad-go/workflows/go%20build/badge.svg
 [build-status-url]: https://github.com/grokify/elastirad-go/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/elastirad-go
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/elastirad-go
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/elastirad-go
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/elastirad-go
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/elastirad-go/blob/master/LICENSE