package requests

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

// ResponseValidator 响应检查器
type ResponseValidator func(*http.Response) error

// Option 参数
type Option func(*Options)

// Options 参数
type Options struct {
	urlValues         url.Values
	header            http.Header
	serializer        Serializer
	httpClient        *http.Client
	responseValidator ResponseValidator
}

var defaultOptions = Options{
	serializer: JSONSerializer(),
	httpClient: &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 15 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	},
	responseValidator: func(resp *http.Response) error {
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected statuscode:%d", resp.StatusCode)
		}
		return nil
	},
}
