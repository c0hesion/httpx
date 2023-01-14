package httpx

import "net/http"

type Response struct {
	Body       any
	StatusCode int
	Headers    http.Header
}

// SetBody sets the body value for the Response.
func (r *Response) SetBody(body any) *Response {
	r.Body = body
	return r
}

// SetStatusCode sets the status code for the Response.
func (r *Response) SetStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

// AddHeader adds a single header to the Response.
func (r *Response) AddHeader(key string, value string) *Response {
	r.Headers.Add(key, value)
	return r
}

// AddHeaders adds multiple headers to the Response.
func (r *Response) AddHeaders(values map[string]string) *Response {
	for k, v := range values {
		r.Headers.Add(k, v)
	}
	return r
}

// NewResponse creates a new Response with the given status and body.
func NewResponse(statusCode int, body any) *Response {
	return &Response{
		Body:       body,
		StatusCode: statusCode,
		Headers:    make(http.Header),
	}
}

// Ok returns a Response initialized with status code 200 and without a body and empty headers.
func Ok() *Response {
	return NewResponse(http.StatusOK, nil)
}

// OkWithBody returns a Response initialized with status code 200 and with a body and empty headers.
func OkWithBody(body any) *Response {
	return NewResponse(http.StatusOK, body)
}

// Created returns a Response initialized with status code 201 and without a body and the Location header set.
func Created(location string) *Response {
	return NewResponse(http.StatusCreated, nil).AddHeader(LocationHeader, location)
}

// Accepted returns a Response initialized with status code 202 and without a body and empty headers.
func Accepted() *Response {
	return NewResponse(http.StatusAccepted, nil)
}

// NoContent returns a Response initialized with status code 204 and without a body and empty headers.
func NoContent() *Response {
	return NewResponse(http.StatusNoContent, nil)
}

// BadRequest returns a Response initialized with status code 400 and without a body and empty headers.
func BadRequest() *Response {
	return NewResponse(http.StatusBadRequest, nil)
}

// NotFound returns a Response initialized with status code 404 and without a body and empty headers.
func NotFound() *Response {
	return NewResponse(http.StatusNotFound, nil)
}

// InternalServerError returns a Response with status code 500 and a body set to the error message.
func InternalServerError(err error) *Response {
	return NewResponse(http.StatusInternalServerError, err.Error())
}

// ResponseOf returns a Response with status code 200 OK and the body set. If err is not nil, the body will be set to
// nil and the status code to 404 Not Found.
//
// This is a useful shortcut for many HTTP related operations.
func ResponseOf(body any, err error) *Response {
	if err != nil {
		return NewResponse(http.StatusNotFound, nil)
	}
	return OkWithBody(body)
}
