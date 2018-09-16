package e2e_helpers

import (
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http"
	"github.com/earcamone/go-mock-yourself/http/helpers"
)

//
// MockResponseMatch() will ensure the received HTTP Response matches the received Mock Response
//

func MockResponseMatch(t *testing.T, response *http.Response, mock *go_mock_yourself_http.Mock) {
	ass := assert.New(t)

	//
	// Get HTTP Response associated request required by Mock methods
	//

	request := response.Request

	//
	// Get HTTP Response Information
	//

	code, headers, body := go_mock_yourself_helpers.ResponseInformation(response)

	//
	// Get Mock Response Information
	//

	expectedBody := mock.Response.GetBody(mock.Name, request)
	expectedHeaders := mock.Response.GetHeaders(mock.Name, request)
	expectedStatusCode := mock.Response.GetStatusCode(mock.Name, request)

	//
	// Ensure HTTP Response status code and body matches Mock Response
	//

	ass.Equal(expectedBody, body, "Request's response body doesn't match Mock Response body. Method: %s, Mock body: %s, response body: %s", request.Method, expectedBody, body)
	ass.Equal(expectedStatusCode, code, "Request's response code doesn't match Mock Response code. Method: %s, Mock code: %d, response code: %d", request.Method, expectedStatusCode, code)

	//
	// Ensure HTTP Response headers match Mock Response headers
	//

	if expectedHeaders != nil {
		ass.Equal(expectedHeaders, headers, "Request's response headers don't match Mock Response headers. Method: %s, Mock headers: %v, response headers: %v", request.Method, expectedHeaders, headers)
	}
}
