package simple

import (
	//"mosn.io/envoy-go-extension/pkg/http/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/api"
)

func ConfigFactory(interface{}) api.StreamFilterFactory {
	return func(callbacks api.FilterCallbackHandler) api.StreamFilter {
		return &filter{
			callbacks: callbacks,
		}
	}
}
