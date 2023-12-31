package elastirad

const (
	// DefaultScheme is the HTTP scheme for the default server.
	DefaultScheme string = "https"
	// DefaultHost is the HTTP host for the default server.
	DefaultHost string = "127.0.0.1:9200"
	// DefaultServerURL is the HTTP host for the default server.
	DefaultServerURL string = "https://127.0.0.1:9200"

	MappingFieldKeyword = "keyword" // see: https://www.elastic.co/blog/strings-are-dead-long-live-strings
	MappingFieldRaw     = "raw"

	// SlugCreate is the URL path part for creates.
	SlugCreate  = "_create"
	SlugDoc     = "_doc"
	SlugMapping = "_mapping"
	// SlugSearch is the URL path part for search.
	SlugSearch = "_search"
	// SlugUpdate is the URL path part for updates.
	SlugUpdate = "_update"

	TypeInteger       = "integer" // More at: https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeKeyword       = "keyword"
	TypeString        = "string" // Deprecation in ES5: https://www.elastic.co/blog/strings-are-dead-long-live-strings
	TypeText          = "text"
	TypeMatchOnlyText = "match_only_text"
)
