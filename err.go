package jrdhttp

import (
	"strconv"
	"strings"
	"net/http"
)

func Err(rw http.ResponseWriter, httpStatusCode uint) {
	_err(rw, httpStatusCode)}

func _err(rw http.ResponseWriter, httpStatusCode uint) {
	if nil == rw {
		return
	}

	var buffer strings.Builder

	buffer.WriteString(`{"__http-response-status-code":"`)
	buffer.WriteString( strconv.FormatUint(uint64(httpStatusCode), 10) )
	buffer.WriteString(`"}`)

	respond(rw, httpStatusCode, buffer.String())
}
