package jrdhttp

import (
	"net/http"
)

func Gone(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusGone

	Err(rw, httpStatusCode)
}
