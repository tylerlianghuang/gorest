package gorest

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// Request struct
type Request struct {
	Body   []byte
	Method string
	Header http.Header
	URL    url.URL
}

func (r Request) String() string {
	s := fmt.Sprintf("Request URL: %s\n", r.URL.String())
	s += fmt.Sprintf("Request Method: %s\n", r.Method)
	s += fmt.Sprintf("Response Body: %s\n", r.Body)

	for k, v := range r.Header {
		s += fmt.Sprintln("Request Header Key: ", k, "Value: ", v)
	}

	return s
}

// Response struct
type Response struct {
	Body       []byte
	Header     http.Header
	StatusCode int
}

func (r Response) String() string {
	s := fmt.Sprintf("Response status code: %d\n", r.StatusCode)
	s += fmt.Sprintf("Response Body: %s\n", string(r.Body))

	for k, v := range r.Header {
		s += fmt.Sprintln("Response Header Key: ", k, "Value: ", v)

	}

	return s
}

// Session struct
type Session struct {
	Cookie *cookiejar.Jar
}