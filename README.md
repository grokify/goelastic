# GoElastic: A Go Client Wrapper for Elasticsearch

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/goelastic/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/goelastic/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/goelastic/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/goelastic/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/goelastic/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/goelastic/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/goelastic
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/goelastic
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/goelastic
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/goelastic
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fgoelastic
 [loc-svg]: https://tokei.rs/b1/github/grokify/goelastic
 [repo-url]: https://github.com/grokify/goelastic
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/goelastic/blob/main/LICENSE

Simple client to query Elasticsearch API using HTTP API documentation.

## Usage

See the sample code in the [docs folder](docs).

So far the following example code has been created:

1. Create Index: [Go code](docs/reference/indices-create-index), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html)
1. Index Docs: [Go code](docs/reference/docs-index), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html)
1. Bool Query: [Go code](docs/reference/query-dsl-bool-query), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html)
1. Terms Aggregation: [Go code](docs/reference/search-aggregations-bucket-terms-aggregation), [ES docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-terms-aggregation.html#search-aggregations-bucket-terms-aggregation)

## References

1. Mapping
    1. [An Introduction to Elasticsearch Mapping](https://www.elastic.co/blog/found-elasticsearch-mapping-introduction)
