package requests

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Requests Requests
type Requests interface {
	Put(ctx context.Context, url string, req interface{}) Result
	Delete(ctx context.Context, url string) Result
	Post(ctx context.Context, url string, req interface{}) Result
	Get(ctx context.Context, url string) Result
}

type requests struct {
	options Options
}

// New New
func New(options ...Option) Requests {
	tmpOptions := defaultOptions
	for _, option := range options {
		option(&tmpOptions)
	}
	return &requests{options: tmpOptions}
}

// Put Put请求
func (r *requests) Put(ctx context.Context, url string, req interface{}) Result {
	return r.do(ctx, http.MethodPut, url, req)

}

// Delete Delete请求
func (r *requests) Delete(ctx context.Context, url string) Result {
	return r.do(ctx, http.MethodDelete, url, nil)
}

// Post Post请求
func (r *requests) Post(ctx context.Context, url string, req interface{}) Result {
	return r.do(ctx, http.MethodPost, url, req)
}

// Get Get请求
func (r *requests) Get(ctx context.Context, url string) Result {
	return r.do(ctx, http.MethodGet, url, nil)
}

func (r *requests) do(ctx context.Context, method, url string, req interface{}) Result {
	if r.options.urlValues != nil {
		url = fmt.Sprintf("%s?%s", url, r.options.urlValues.Encode())
	}

	httpRequest, err := http.NewRequest(method, url, nil)
	if err != nil {
		return &result{err: err}
	}

	for key, values := range r.options.header {
		for _, value := range values {
			httpRequest.Header.Add(key, value)
		}
	}

	if req != nil {
		buf := bytes.NewBuffer(nil)
		if err := r.options.serializer.Marshal(buf)(req); err != nil {
			return &result{err: err}
		}
		httpRequest.Body = ioutil.NopCloser(buf)
	}
	httpRequest = httpRequest.WithContext(ctx)

	resp, err := r.options.httpClient.Do(httpRequest)
	if err != nil {
		return &result{err: err}
	}

	if err := r.options.responseValidator(resp); err != nil {
		resp.Body.Close()
		return &result{err: err}
	}

	return &result{rc: resp.Body, decode: r.options.serializer.Unmarshal}
}
