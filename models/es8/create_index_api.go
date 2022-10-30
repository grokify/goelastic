package es8

// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html#indices-create-index
// https://www.elastic.co/blog/strings-are-dead-long-live-strings
// https://discuss.elastic.co/t/index-not-analyzed/126606

const (
	DefaultBaseURL = "https://localhost:9200"
	TypeInteger    = "integer"
	TypeKeyword    = "keyword"
	TypeText       = "text"
)

type CreateIndexBody struct {
	Settings *Settings `json:"settings,omitempty"`
	Mappings *Mappings `json:"mappings,omitempty"`
}

type Settings struct {
	Index            *Index `json:"index,omitempty"`
	NumberOfShards   uint32 `json:"number_of_shards,omitempty"`
	NumberOfReplicas uint32 `json:"number_of_replicas,omitempty"`
}

type Index struct {
	NumberOfShards   uint32 `json:"number_of_shards,omitempty"`
	NumberOfReplicas uint32 `json:"number_of_replicas,omitempty"`
}

type Mappings struct {
	// All        All                 `json:"_all,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
}

type All struct {
	Enabled bool `json:"enabled,omitempty"`
}

type Property struct {
	Type        string              `json:"type,omitempty"`
	Index       bool                `json:"index"`
	Format      string              `json:"format,omitempty"`
	Path        string              `json:"path,omitempty"`
	IgnoreAbove int                 `json:"ignore_above,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
	Fields      map[string]Property `json:"fields,omitempty"` // key can be "raw"
}

func SettingsTest() *Settings {
	return &Settings{
		NumberOfShards: 1,
	}
}
