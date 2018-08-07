package go_mock_yourself_request_test

import (
	"testing"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http"
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
)

//
// TestRequestsHeadersRegexMatching() will ensure Requests Headers regular expressions matching is working properly
//

func TestRequestsHeadersRegexMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Set always matching Headers regular expression
	//

	matchingHeaders := make(map[string]string)
	matchingHeaders[".*"] = ".*"

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetHeaders(matchingHeaders)

	for i := 0; i < 100; i++ {
		randomString := go_mock_yourself_helpers.RandomString(100)

		randomHeaders := make(map[string]string)
		randomHeaders[randomString] = randomString

		ass.True(mockRequest.HeadersMatch(randomHeaders))
	}

	//
	// Set never matching Headers regular expression
	//

	matchingHeaders = make(map[string]string)
	matchingHeaders["this header"] = "will never match"

	mockRequest.SetHeaders(matchingHeaders)

	for i := 0; i < 100; i++ {
		randomString := go_mock_yourself_helpers.RandomString(100)

		randomHeaders := make(map[string]string)
		randomHeaders[randomString] = randomString

		ass.False(mockRequest.HeadersMatch(randomHeaders))
	}

	//
	// Set specific Headers regular expressions matching
	//

	matchingHeaders = make(map[string]string)
	matchingHeaders["this .* header"] = "should .* match"

	mockRequest.SetHeaders(matchingHeaders)

	requestHeaders := make(map[string]string)
	requestHeaders["the following header"] = "will match"
	requestHeaders["this wefn2309fj header"] = "should 2q39rjfw2ogf match"

	ass.True(mockRequest.HeadersMatch(requestHeaders))

	//
	// Ensure ALL specified Headers matching is working properly
	//

	matchingHeaders = make(map[string]string)
	matchingHeaders["this .* header"] = "should .* match"
	matchingHeaders["this matching header won't match"] = "<--"

	mockRequest.SetHeaders(matchingHeaders)

	requestHeaders = make(map[string]string)
	requestHeaders["the header"] = "wont match"
	requestHeaders["this wefn2309fj header"] = "should 2q39rjfw2ogf match"

	ass.False(mockRequest.HeadersMatch(requestHeaders))

	//
	// Ensure multiple specified Headers matching is working properly
	//

	matchingHeaders = make(map[string]string)
	matchingHeaders["this .* header"] = "should .* match"
	matchingHeaders["this static header"] = "will match"

	mockRequest.SetHeaders(matchingHeaders)

	requestHeaders = make(map[string]string)
	requestHeaders["this header"] = "wont match"
	requestHeaders["this static header"] = "will match"
	requestHeaders["this wefn2309fj header"] = "should 2q39rjfw2ogf match"

	ass.True(mockRequest.HeadersMatch(requestHeaders))
}

//
// TestRequestsHeadersCallbackMatching() will ensure Requests Headers Matching Callback scheme is working properly
//

func TestRequestsHeadersCallbackMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Generate random Headers to assert inside Request Matching Callback its being called properly
	//

	randomString1 := go_mock_yourself_helpers.RandomString(100)
	randomString2 := go_mock_yourself_helpers.RandomString(100)

	//
	// Create Request Headers
	//

	requestHeaders := make(map[string]string)
	requestHeaders[randomString1] = randomString2
	requestHeaders[randomString2] = randomString1

	//
	// Create Request
	//

	mockRequest := new(go_mock_yourself_http.Request)

	//
	// Ensure Request Matching Callback scheme works properly
	//

	matchingCallbackResults := []bool { true, false }

	for _, matchingResult := range matchingCallbackResults {
		//
		// Set Matching Callback based on current bool value loop
		//

		mockRequest.SetHeaders(func(requestHeaders map[string]string) bool {
			ass.Equal(requestHeaders, requestHeaders)
			return matchingResult
		})

		//
		// Set always matching static regular expression which should be ignored as there is a Matching Callback registered
		//

		alwaysMatchingHeaders := make(map[string]string)
		alwaysMatchingHeaders[".*"] = ".*"

		mockRequest.SetHeaders(alwaysMatchingHeaders)

		//
		// Matching Callback criteria should be used ignoring always matching static regular expression
		//

		ass.Equal(matchingResult, mockRequest.HeadersMatch(requestHeaders))
	}
}

//
// TestRequestsHeadersMissingCriteriaMatching() will ensure Requests Headers missing criteria is working properly
//

func TestRequestsHeadersMissingCriteriaMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Generate random Headers to assert inside Request Matching Callback its being called properly
	//

	randomString1 := go_mock_yourself_helpers.RandomString(100)
	randomString2 := go_mock_yourself_helpers.RandomString(100)

	//
	// Create Request Headers
	//

	requestHeaders := make(map[string]string)
	requestHeaders[randomString1] = randomString2
	requestHeaders[randomString2] = randomString1

	//
	// Create Request
	//

	mockRequest := new(go_mock_yourself_http.Request)

	//
	// Ensure Request matches any URL on missing URL matching criteria
	//

	ass.True(mockRequest.HeadersMatch(requestHeaders))
}
