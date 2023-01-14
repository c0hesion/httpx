package httpx

import "net/http"

type Request[T any] struct {
	http.Request
	Body T
}
