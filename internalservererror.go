package jrdhttp

import (
	"net/http"
)

func InternalServerError(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusInternalServerError

	Err(rw, httpStatusCode)
}
