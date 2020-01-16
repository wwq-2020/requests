package requests

import (
	"encoding/json"
	"io"
)

// Serializer 序列化
type Serializer interface {
	Marshal(io.Writer) func(interface{}) error
	Unmarshal(io.Reader) func(interface{}) error
}

type jsonSerializer struct {
}

// JSONSerializer json序列化
func JSONSerializer() Serializer {
	return &jsonSerializer{}
}

func (s *jsonSerializer) Marshal(w io.Writer) func(interface{}) error {
	return json.NewEncoder(w).Encode
}

func (s *jsonSerializer) Unmarshal(r io.Reader) func(interface{}) error {
	return json.NewDecoder(r).Decode
}
