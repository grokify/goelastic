package main

import (
	"fmt"

	"github.com/grokify/elastirad-go/models/es8/es8examples"
	"github.com/grokify/goauth"
	"github.com/grokify/gohttp/httpsimple"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/log/logutil"
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

	hclient, err := goauth.NewClientBasicAuth(opts.Username, opts.Password, true)
	logutil.FatalErr(err)

	sclient := httpsimple.SimpleClient{
		BaseURL:    "https://localhost:9200",
		HTTPClient: hclient}

	sreq := es8examples.ExampleCreateIndexMappings()

	resp, err := sclient.Do(*sreq)
	logutil.FatalErr(err)

	data, err := jsonutil.PrettyPrintReader(resp.Body, "", "  ")
	logutil.FatalErr(err)

	fmt.Println(string(data))

	fmt.Println("DONE")
}
