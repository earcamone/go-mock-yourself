package end_to_end_test

import (
	"testing"
	"strings"
	"net/http"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http"
	"github.com/mercadolibre/go-mock-yourself/http/helpers"

	// Go Mock Yourself e2e Tests Imports
	"github.com/mercadolibre/go-mock-yourself/http/tests/internal/e2e_server"
	"github.com/mercadolibre/go-mock-yourself/http/tests/internal/e2e_helpers"
)

//
// e2e tests will run in the background a "Testing HTTP Server" which will return for any supported HTTP method request
// a random status code with random headers and body.
//
// All test cases will then create a Go Mock Yourself Mock using the specific scheme we want to test together with a
// matching HTTP request for it, then the test case will call mockShouldMatch() which will:
//
//  1. Send the received HTTP request to the Testing HTTP Server and ensure it receives its response
//  2. Install the received Mock and re-send the request, this time it will ensure it received the Mock Response
//  3. It will Pause() Go Mock Yourself and re-send the request expecting to receive the Testing HTTP Server response
//  4. It will Play() Go Mock Yourself and re-send the request expecting to receive again the Mock Response
//  5. Finally it will alter mock setting an always false ShouldMock() callback expecting to redirect the matching
//     HTTP request to the Testing HTTP server instead of returning Mock Response.
//

func init() {
	//
	// Start end-to-end Testing HTTP Server
	//

	e2e_server.Run()
}

//
// mockShouldMatch() will ensure the received Mock works properly, additionally
// testing generic Mock schemes such as Play/Pause and ShouldMock schemes.
//

func mockShouldMatch(t *testing.T, request *http.Request, mySweetMock *go_mock_yourself_http.Mock, mockShouldNotMatch ...bool) {
	mockShouldMatch := len(mockShouldNotMatch) == 0 || mockShouldNotMatch[0] == false

	//
	// Reset Go Mock Yourself Mocks
	//

	go_mock_yourself_http.Reset()

	//
	// Do() will empty Request body stream internally which is required by subsequent Request body comparisons,
	// thus we will duplicate the received request to fed Do() with a valid request with the corresponding body
	//

	requestCopy1 := duplicateRequest(request)
	requestCopy2 := duplicateRequest(request)
	requestCopy3 := duplicateRequest(request)
	requestCopy4 := duplicateRequest(request)
	requestCopy5 := duplicateRequest(request)

	//
	// Send request to our listening HTTP Testing Server through Go's native http package
	//

	response, _ := http.DefaultClient.Do(request)

	//
	// Ensure response matches HTTP Testing Server response (request was processed by it)
	//

	e2e_server.ResponseMatch(t, response)

	//
	// Install Always Matching Mock so it handles our previous request when we re-attempt it
	//

	go_mock_yourself_http.Install(mySweetMock)

	//
	// Send request to our listening HTTP Testing Server through Go's native http package
	//

	response, _ = http.DefaultClient.Do(requestCopy1)

	//
	// Ensure Response matches either Mock Response or e2e Testing Server Response depending on the received Mock Request Match
	//

	switch mockShouldMatch {
	case false:
		e2e_server.ResponseMatch(t, response)

	case true:
		e2e_helpers.MockResponseMatch(t, response, mySweetMock)
	}

	//
	// Pause Go Mock Yourself Scheme (now requests should again be redirected to the HTTP Testing Server)
	//

	go_mock_yourself_http.Pause()

	//
	// Send request to our listening HTTP Testing Server through Go's native http package
	//

	response, _ = http.DefaultClient.Do(requestCopy2)

	//
	// Ensure response matches HTTP Testing Server response
	//

	e2e_server.ResponseMatch(t, response)

	//
	// Play Go Mock Yourself Scheme
	//

	go_mock_yourself_http.Play()

	//
	// Send request to our listening HTTP Testing Server through Go's native http package
	//

	response, _ = http.DefaultClient.Do(requestCopy3)

	//
	// Ensure Response matches Mock Response
	//

	switch mockShouldMatch {
	case false:
		e2e_server.ResponseMatch(t, response)

	case true:
		e2e_helpers.MockResponseMatch(t, response, mySweetMock)
	}

	//
	// Reset Go Mock Yourself Mocks
	//

	go_mock_yourself_http.Reset()

	//
	// Remove Mock Response, Mock should match BUT still redirect to native Do() method. This feature is usually used
	// just to log specific requests flows, alter their processing time with the Timeout option or many other cool
	// stuff you can do with your imagination :P
	//

	mockResponse := mySweetMock.Response
	mySweetMock.Response = nil

	//
	//
	//

	go_mock_yourself_http.Install(mySweetMock)

	//
	// Send request to our listening HTTP Testing Server through Go's native http package
	//

	response, _ = http.DefaultClient.Do(requestCopy4)

	//
	// Mock should had ignored current request due to just installed "always fail" ShouldMock callback
	//

	e2e_server.ResponseMatch(t, response)

	//
	// Reset Go Mock Yourself Mocks
	//

	go_mock_yourself_http.Reset()

	//
	//
	//

	mySweetMock.Response = mockResponse

	//
	// Set Mock's ShouldMock scheme to force non matching criteria
	//

	mySweetMock.Request.SetShouldMock(func(*http.Request) bool {
		return false
	})

	//
	//
	//

	go_mock_yourself_http.Install(mySweetMock)

	//
	// We already tested that when installed properly the received Mock processes the request,
	// let's now ensure the ShouldMock scheme is working properly by forcing request ignorance
	//

	response, _ = http.DefaultClient.Do(requestCopy5)

	//
	// Mock should had ignored current request due to just installed "always fail" ShouldMock callback
	//

	e2e_server.ResponseMatch(t, response)
}

func mockShouldNotMatch(t *testing.T, request *http.Request, mySweetMock *go_mock_yourself_http.Mock) {
	mockShouldMatch(t, request, mySweetMock, true)
}

func duplicateRequest(request *http.Request) *http.Request {
	requestBody, _ := go_mock_yourself_helpers.GetBody(request)
	requestCopy, _ := http.NewRequest(request.Method, request.URL.String(), strings.NewReader(string(requestBody)))

	if len(request.Header) > 0 {
		headers := go_mock_yourself_helpers.HTTPHeadersToMap(request.Header)
		requestCopy.Header = go_mock_yourself_helpers.HeadersMapToHTTPHeaders(headers)
	}

	return requestCopy
}
