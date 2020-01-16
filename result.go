package requests

import "io"

// Result 响应
type Result interface {
	Unmarshal(interface{}) error
}

// result 响应
type result struct {
	rc     io.ReadCloser
	err    error
	decode func(io.Reader) func(interface{}) error
}

// Unmarshal 反序列化
func (r *result) Unmarshal(resp interface{}) error {
	defer r.rc.Close()
	if r.err != nil {
		return r.err
	}
	if err := r.decode(r.rc)(resp); err != nil {
		return err
	}
	return nil
}
