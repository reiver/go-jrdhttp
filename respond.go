package jrdhttp

import (
	"io"
	"net/http"
)

func respond(rw http.ResponseWriter, httpStatusCode uint, content string) {

	// Specify the Content-Type of the HTTP response.
	rw.Header().Set("Content-Type", "application/jrd+json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	// Write the HTTP headers in the response, with the specified HTTP response code.
	rw.WriteHeader(int(httpStatusCode))

	// Write the HTTP response content.
	io.WriteString(rw, content)
}


