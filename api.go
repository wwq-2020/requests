package requests

import (
	"context"
)

// Put Put请求
func Put(ctx context.Context, url string, req interface{}, options ...Option) Result {
	return New(options...).Put(ctx, url, req)
}

// Delete Delete请求
func Delete(ctx context.Context, url string, options ...Option) Result {
	return New(options...).Delete(ctx, url)
}

// Post Post请求
func Post(ctx context.Context, url string, req interface{}, options ...Option) Result {
	return New(options...).Post(ctx, url, req)
}

// Get Get请求
func Get(ctx context.Context, url string, options ...Option) Result {
	return New(options...).Get(ctx, url)
}
