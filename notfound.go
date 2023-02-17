package jrdhttp

import (
	"net/http"
)

func NotFound(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusNotFound

	Err(rw, httpStatusCode)
}
