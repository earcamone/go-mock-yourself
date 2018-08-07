package go_mock_yourself_http

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
)

func createHttpResponseFromMock(request *http.Request, mock Mock) (*http.Response, error) {
	//
	// Mock Response should return an error?
	//

	errorResponse := mock.Response.GetError(mock.Name, request)

	if errorResponse != nil {
		return nil, errorResponse
	}

	//
	// Get Mock Response Status Code
	//

	statusCode := mock.Response.GetStatusCode(mock.Name, request)

	//
	// Get Mock Response body and build the corresponding stream required by http.Response
	//

	responseBody := mock.Response.GetBody(mock.Name, request)

	bodyBuffer := bytes.NewBuffer([]byte(responseBody))
	bodyStream := ioutil.NopCloser(bodyBuffer)

	//
	// Get Mock Response Headers
	//

	headers := mock.Response.GetHeaders(mock.Name, request)

	//
	// Build Mock native HTTP Response
	//

	response := http.Response {
		//
		//
		//

		StatusCode: statusCode,

		//
		//
		//

		Header: go_mock_yourself_helpers.HeadersMapToHTTPHeaders(headers),

		//
		//
		//

		Body: bodyStream,

		//
		// ===========================================================================================
		//   Here on you will find just a quick attempt to mimik common http.Response objects values
		// ===========================================================================================
		//

		//
		// Original http.Request that generated this response
		//

		Request: request,

		//
		// Is content uncompressed? come on! are you really going to serve a mock with compressed content?!?
		// if after taking the appropriate time you thought about it with an old-fashion in hand and you still
		// think is mandatory for you to serve the content compressed send me a message so i enable this flag
		// setting for you in a nice way.
		//

		Uncompressed: true,

		//
		// According to documentation is the number of bytes that can be safely read from body. Even though
		// we can easily calculate this from the Mock response body, while performing test with different http
		// package helper functions (http.Get, http.Post) the value they return on http.Response objects seem
		// always to be -1 so lets just copy that behaviour.
		//

		ContentLength: -1,

		//
		// bleh..
		//

		Status:  fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)),
		Proto:   "HTTP/1.1",

		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	return &response, nil
}
