package jrdhttp

import (
	"net/http"
)

func Conflict(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusConflict

	Err(rw, httpStatusCode)
}
