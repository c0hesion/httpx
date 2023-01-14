package httpx

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc[T any] func(request *http.Request, body *T) *Response

func Handler[RequestBody any, ResponseBody any](handlerFunc HandlerFunc[RequestBody]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			body               = new(RequestBody)
			response *Response = nil
		)
		// decode the body and call handler func
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			response = handlerFunc(r, nil)
		} else {
			response = handlerFunc(r, body)
		}
		// this should never happen, and if it does, we should return 500
		if response == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(response.StatusCode)
		_, _ = w.Write(transformBody[ResponseBody](response.Body))
	}
}

func transformBody[T any](body any) []byte {
	switch v := body.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	case T:
		if b, err := json.Marshal(v); err == nil {
			return b
		}
	}
	return []byte{}
}
