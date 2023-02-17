package jrdhttp

import (
	"net/http"
)

func BadRequest(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusBadRequest

	Err(rw, httpStatusCode)
}
