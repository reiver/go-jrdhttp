package jrdhttp

import (
	"net/http"
)

func NotImplemented(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusNotImplemented

	Err(rw, httpStatusCode)
}
