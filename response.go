package httpx

import "net/http"

type Response[T any] struct {
	Body       T
	StatusCode int
	Headers    http.Header
}

func (r *Response[T]) SetBody(body T) *Response[T] {
	r.Body = body
	return r
}

func (r *Response[T]) SetStatusCode(statusCode int) *Response[T] {
	r.StatusCode = statusCode
	return r
}

func (r *Response[T]) AddHeader(key string, value string) *Response[T] {
	r.Headers.Add(key, value)
	return r
}

func (r *Response[T]) AddHeaders(keyValues ...string) *Response[T] {
	if len(keyValues)%2 != 0 {
		panic("keyValues is not divisible by 2")
	}
	for i := 0; i < len(keyValues)/2; i++ {
		r.Headers.Add(keyValues[i*2], keyValues[i*2+1])
	}
	return r
}

func NewResponse[T any](statusCode int, body T) *Response[T] {
	return &Response[T]{
		Body:       body,
		StatusCode: statusCode,
		Headers:    make(http.Header),
	}
}

// Ok returns a Response initialized with status code 200 and without a body and empty headers.
func Ok() *Response[any] {
	return NewResponse[any](http.StatusOK, nil)
}

// OkWithBody returns a Response initialized with status code 200 and with a body and empty headers.
func OkWithBody[T any](body T) *Response[T] {
	return NewResponse(http.StatusOK, body)
}

// Created returns a Response initialized with status code 201 and without a body and the Location header set.
func Created(location string) *Response[any] {
	return NewResponse[any](http.StatusCreated, nil).AddHeader(LocationHeader, location)
}

// Accepted returns a Response initialized with status code 202 and without a body and empty headers.
func Accepted() *Response[any] {
	return NewResponse[any](http.StatusAccepted, nil)
}

// NoContent returns a Response initialized with status code 204 and without a body and empty headers.
func NoContent() *Response[any] {
	return NewResponse[any](http.StatusNoContent, nil)
}

// BadRequest returns a Response initialized with status code 400 and without a body and empty headers.
func BadRequest() *Response[any] {
	return NewResponse[any](http.StatusBadRequest, nil)
}

// NotFound returns a Response initialized with status code 404 and without a body and empty headers.
func NotFound() *Response[any] {
	return NewResponse[any](http.StatusNotFound, nil)
}

// InternalServerError returns a Response with status code 500 and a body set to the error message.
func InternalServerError(err error) *Response[string] {
	return NewResponse(http.StatusInternalServerError, err.Error())
}

// ResponseOf returns a Response with status code 200 OK and the body set. If err is not nil, the body will be set to
// nil and the status code to 404 Not Found.
//
// This is a useful shortcut for many HTTP related operations.
func ResponseOf[T any](body T, err error) *Response[T] {
	if err != nil {
		return NewResponse[T](http.StatusNotFound, nil)
	}
	return OkWithBody(body)
}
