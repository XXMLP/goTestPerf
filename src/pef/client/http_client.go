package client

import (
	"time"

	"github.com/valyala/fasthttp"
)

var (
	// HTTPClient global http client object
	HTTPClient *fasthttp.Client = &fasthttp.Client{
		MaxConnsPerHost: 16384, // MaxConnsPerHost  default is 512, increase to 16384
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
	}
)
