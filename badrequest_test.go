package jrdhttp_test

import (
	"github.com/reiver/go-jrdhttp"

	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestBadRequest(t *testing.T) {

	const expectedStatusCode = http.StatusBadRequest

	var rr *httptest.ResponseRecorder = httptest.NewRecorder()
	var rw http.ResponseWriter = rr

	jrdhttp.BadRequest(rw)

	var response *http.Response
	{
		response = rr.Result()
		if nil == response {
			t.Error("nil response")
			return
		}
	}

	{
		var expected int = expectedStatusCode
		var actual   int = response.StatusCode

		if expected != actual {
			t.Error("the actual HTTP response status-code is not what was expected")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	var contentType string
	{
		var headers http.Header = rr.Header()
		if nil == headers {
			t.Error("did not expect headers to be nil, but actually was")
			return
		}

		contentTypes, found := headers["Content-Type"]
		if !found {
			t.Error("Content-Type header found")
			return
		}

		{
			const expected int = 1
			var   actual   int = len(contentTypes)

		if expected != actual {
			t.Error("The actual number of Content-Type headers is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
		}

		contentType = contentTypes[0]
	}

	{
		const expected string = "application/jrd+json"
		var   actual   string = contentType

		if expected != actual {
			t.Error("The actual value for Content-Type is not what was expected")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}

	body, err := io.ReadAll(response.Body)
	if nil != err {
		t.Error("did not expect to get an error but actually got one")
		t.Logf("ERROR: (%T) %q", err, err)
		return
	}

	{
		var expected string = fmt.Sprintf(`{"__http-response-status-code":"%d"}`, expectedStatusCode)
		var actual   string = string(body)

		if expected != actual {
			t.Error("The actual JRD value is not what was expected")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}

	}
}
