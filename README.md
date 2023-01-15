httpx
=====

`httpx` is a Go library that seeks to extend the `net/http` standard library with some useful types and convenience
functions to reduce the amount boilerplate necessary for common HTTP related tasks and development. The library was
greatly influenced by the design of [James Moiron's `sqlx`](https://github.com/jmoiron/sqlx) library that provides
similar functionality to the `database/sql` standard library.

usage
-----

The primary use case for `httpx` are the `httpx.Response` and `httpx.Request` objects, which wrap and extend
`http.ResponseWriter` and `http.Request` respectively, which together with the `httpx.Handler` function can greatly
reduce the amount of boilerplate required for writing HTTP handlers.

The `httpx.Response` object simply provides a simple way to set headers, a body and status code. This is especially
expressive when combined with the initializing convenience functions. The library has absolutely no dependencies and 
should play well with anything that interfaces with `net/http`.

The following example highlights some common use cases of `httpx`.

```go
package main

import (
	"fmt"
	"github.com/c0hesion/httpx"
	"net/http"
)
import "github.com/go-chi/chi/v5"

type HelloRequest struct {
	Name string `json:"name"`
}

// Define HTTP handlers that take in Request and produce Response.
func helloHandler(request *httpx.Request[HelloRequest]) *httpx.Response {
	// Request can automatically parse JSON objects within.
	body, err := request.JSON()
	if err != nil {
		// Easily return common responses
		return httpx.BadRequest()
	}
	hello := fmt.Sprintf("hello, %s!", body.Name)
	return httpx.OkWithBody(hello)
}

func teapotHandler(request *httpx.Request[any]) *httpx.Response {
	// Of course, you can easily form your own responses too
	return httpx.NewResponse(
		http.StatusTeapot, 
		"Are you looking for coffee.com?", 
	)
	// Naturally, you can return any JSON struct, byte array or string.
}

func cookieHandler(request *httpx.Request[HelloRequest]) *httpx.Response {
	// httpx.Request is just an extension of http.Request, so all methods are available
	cookie, err := request.Cookie("example")
	if err != nil {
		reqBody, _ := request.JSON()
		var cookie = http.Cookie{
			Name:     "example",
			Value:    fmt.Sprintf("hello, %s", reqBody.Name),
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		// Easily add cookies with the Response object
		return httpx.Ok().AddCookie(&cookie)
	} else {
		return httpx.OkWithBody(cookie.Value)
	}
}

func main() {
	r := chi.NewRouter()
	// Handler can wrap any httpx.HandlerFunc as a http.HandlerFunc
	r.Get("/hello", httpx.Handler[string](helloHandler))
	r.Get("/teapot", httpx.Handler[any](teapotHandler))
	r.Post("/cookie", httpx.Handler[any](cookieHandler))
}
```

contributing
------------

`httpx` is still early in development, you can contribute by checking the issues tracker.