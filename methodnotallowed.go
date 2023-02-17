package jrdhttp

import (
	"net/http"
)

func MethodNotAllowed(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusMethodNotAllowed

	Err(rw, httpStatusCode)
}
