package reference

import (
	"fmt"
	"net/http"

	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/log/logutil"
)

const (
	Index        = "twitter"
	Type         = "tweet"
	IndexAndType = "twitter/tweet"
)

// Example from:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html

// Tweet is an example Twitter Tweet struct.
type Tweet struct {
	User     User     `json:"user,omitempty"`
	PostDate string   `json:"post_date,omitempty"`
	Message  string   `json:"message,omitempty"`
	HashTags []string `json:"hash_tags,omitempty"`
}

type User struct {
	Username   string   `json:"username,omitempty"`
	Location   Location `json:"location,omitempty"`
	IsVerified bool     `json:"is_verified,omitempty"`
}

type Location struct {
	DisplayName string  `json:"display_name,omitempty"`
	Locode      string  `json:"locode,omitempty"`
	Lat         float64 `json:"lat,omitempty"`
	Lon         float64 `json:"lon,omitempty"`
}

func ProcResponse(resp *http.Response, err error) {
	logutil.FatalErr(err)
	if resp == nil {
		return
	}
	body, err := jsonutil.IndentReader(resp.Body, "", "  ")
	logutil.FatalErr(err)
	fmt.Printf("C_RES_BODY: %v\n", string(body))
	fmt.Printf("C_RES_STATUS: %v\n", resp.StatusCode)
}
