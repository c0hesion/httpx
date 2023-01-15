package httpx

import (
	"encoding/json"
	"net/http"
)

type Request[T any] struct {
	*http.Request
}

// JSON attempts to decode the request body as JSON.
func (r *Request[T]) JSON() (*T, error) {
	body := new(T)
	err := json.NewDecoder(r.Body).Decode(&body)
	return body, err
}
