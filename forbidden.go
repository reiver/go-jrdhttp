package jrdhttp

import (
	"net/http"
)

func Forbidden(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusForbidden

	Err(rw, httpStatusCode)
}
