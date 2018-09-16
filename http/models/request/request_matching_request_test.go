package go_mock_yourself_request

import (
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"
)

//
// TestRequestsRequestMatchingCallback() will ensure SetRequestCriteria() scheme is working correctly
//

func TestRequestsRequestMatchingCallback(t *testing.T) {
	ass := assert.New(t)

	//
	// Create Request
	//

	mockRequest := new(Request)

	//
	// Ensure Request Matching Callback scheme works properly
	//

	matchingCallbackResults := []bool { true, false }

	for _, matchingResult := range matchingCallbackResults {
		myIncomingRequest := new(http.Request)

		//
		// Set Matching Callback based on current bool value loop
		//

		mockRequest.SetRequestCriteria(func(request *http.Request) bool {
			ass.Equal(myIncomingRequest, request)
			return matchingResult
		})

		//
		// Set always matching dynamic callbacks for all available criterias, they should never be called as if there is
		// a "Request Matching Callback" registered, its return value will determine if a request is or not a match,
		// without even attempting to apply the static/dynamic specific criterias (url, headers, body, etc).
		//

		callbacksCounter := 0

		mockRequest.SetMethod(func(string) bool {
			callbacksCounter++
			return false
		})

		mockRequest.SetUrl(func(string) bool {
			callbacksCounter++
			return false
		})

		mockRequest.SetHeaders(func(map[string]string) bool {
			callbacksCounter++
			return false
		})

		mockRequest.SetBody(func(string) bool {
			callbacksCounter++
			return false
		})

		//
		// Matching Request Callback criteria should determine incoming HTTP
		// request matching, ignoring all other registered Matching Callbacks
		//

		ass.Equal(matchingResult, mockRequest.Match(myIncomingRequest))

		//
		// Counter should be 0 as all other Matching Callbacks should not had been called
		//

		ass.Zero(callbacksCounter)
	}
}

//
// TestRequestsShouldMock() will ensure SetShouldMock() scheme is working correctly
//

func TestRequestsShouldMock(t *testing.T) {
	ass := assert.New(t)

	//
	// Create Request
	//

	mockRequest := new(Request)

	//
	// Ensure Request Matching Callback scheme works properly
	//

	matchingCallbackResults := []bool { true, false }

	for _, matchingResult := range matchingCallbackResults {
		myIncomingRequest := new(http.Request)

		//
		// Lets make Mock Request always match
		//

		mockRequest.SetRequestCriteria(func(request *http.Request) bool {
			return true
		})

		//
		// Let's now install a ShouldMock() function that will conditionally fail based on current loop value
		//

		mockRequest.SetShouldMock(func(*http.Request) bool {
			return matchingResult
		})

		//
		// Always matching Mock Request should only match when Should Mock function returns true
		//

		ass.Equal(matchingResult, mockRequest.ShouldMock(myIncomingRequest))
	}
}
