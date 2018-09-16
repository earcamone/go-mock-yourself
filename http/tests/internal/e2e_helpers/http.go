package e2e_helpers

import (
	"strings"
	"net/http"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http/helpers"
)

const testingHttpServerUrl = "http://127.0.0.1:8080"

//
// createHTTPRequest() will build a Go's http package request
//

func CreateHTTPRequest(method string, url string, headers map[string]string, body string) *http.Request {
	e2eServerUrl := testingHttpServerUrl + url
	request, _ := http.NewRequest(method, e2eServerUrl, strings.NewReader(body))

	if len(headers) > 0 {
		request.Header = go_mock_yourself_helpers.HeadersMapToHTTPHeaders(headers)
	}

	return request
}
