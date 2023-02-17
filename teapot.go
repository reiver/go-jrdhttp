package jrdhttp

import (
	"net/http"
)

func Teapot(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusTeapot

	Err(rw, httpStatusCode)
}
