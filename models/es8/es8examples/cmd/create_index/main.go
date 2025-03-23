package main

import (
	"fmt"

	"github.com/grokify/goauth/authutil"
	"github.com/grokify/goelastic"
	"github.com/grokify/goelastic/models/es8/es8examples"
	"github.com/grokify/mogo/encoding/jsonutil/jsonraw"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Username string `short:"u" long:"username" description:"Basic Auth Username"`
	Password string `short:"p" long:"password" description:"Basic Auth Password"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	logutil.FatalErr(err)

	hclient, err := authutil.NewClientBasicAuth(opts.Username, opts.Password, true)
	logutil.FatalErr(err)

	sclient := httpsimple.Client{
		BaseURL:    goelastic.DefaultServerURL,
		HTTPClient: hclient}

	sreq := es8examples.ExampleCreateIndexMappings()

	resp, err := sclient.Do(*sreq)
	logutil.FatalErr(err)

	data, err := jsonraw.Indent(resp.Body, "", "  ")
	logutil.FatalErr(err)

	fmt.Println(string(data))

	fmt.Println("DONE")
}
