package main

import (
	"envoy-go-filter-samples/simple"

	//"mosn.io/envoy-go-extension/pkg/http"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
)

func init() {
	http.RegisterHttpFilterConfigFactory("envoy-go-test-example", simple.ConfigFactory)
}

func main() {
}
