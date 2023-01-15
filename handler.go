package httpx

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc[T any] func(request *Request[T]) *Response

// Handler creates a http.HandlerFunc which wraps a HandlerFunc in some common boilerplate.
func Handler[Res any, Req any](handlerFunc HandlerFunc[Req]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &Request[Req]{
			Request: r,
		}
		response := handlerFunc(request)
		// this should never happen, and if it does, we should return 500
		if response == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(response.StatusCode)
		_, _ = w.Write(transformBodyToBytes[Res](response.Body))
	}
}

// transformBodyToBytes will transform body to a byte array.
//
// Internally this function type-checks body. A byte array will simply be returned, a string will be turned into bytes
// and body is of type T or a pointer to T, it will be encoded as JSON. Any other value is considered an invalid type
// and will be return an empty byte array.
func transformBodyToBytes[T any](body any) []byte {
	switch v := body.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	case T, *T:
		if b, err := json.Marshal(v); err == nil {
			return b
		}
	}
	return []byte{}
}
