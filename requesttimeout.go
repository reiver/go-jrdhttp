package jrdhttp

import (
	"net/http"
)

func RequestTimeout(rw http.ResponseWriter) {

	const httpStatusCode uint = http.StatusRequestTimeout

	Err(rw, httpStatusCode)
}
