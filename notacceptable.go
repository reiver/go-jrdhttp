package jrdhttp

import (
	"net/http"
)

func NotAcceptable(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusNotAcceptable

	Err(rw, httpStatusCode)
}
