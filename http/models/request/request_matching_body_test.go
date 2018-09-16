package go_mock_yourself_request

import (
	"testing"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http"
	"github.com/earcamone/go-mock-yourself/http/helpers"
)

//
// TestRequestsBodyRegexMatching() will ensure Requests Body regular expressions matching is working properly
//

func TestRequestsBodyRegexMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Set always matching Body regular expression
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetBody(".*")

	for i := 0; i < 100; i++ {
		randomBody := go_mock_yourself_helpers.RandomString(100)
		ass.True(mockRequest.BodyMatch(randomBody))
	}

	//
	// Set never matching Body regular expression
	//

	mockRequest.SetBody("this body will never match, unlike the love we feel with mumi, awwwww")

	for i := 0; i < 100; i++ {
		randomBody := go_mock_yourself_helpers.RandomString(100)
		ass.False(mockRequest.BodyMatch(randomBody))
	}

	//
	// Set specific Body regular expressions matching
	//

	mockRequest.SetBody("this body should .* hopefully match")
	ass.True(mockRequest.BodyMatch("this body should fo9n2193f02efj hopefully match"))
}

//
// TestRequestsBodyCallbackMatching() will ensure Requests Body Matching Callback scheme is working properly
//

func TestRequestsBodyCallbackMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Generate random URL to assert inside Request Matching Callback its being called properly
	//

	randomBody := go_mock_yourself_helpers.RandomString(100)

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

		mockRequest.SetBody(func(requestBody string) bool {
			ass.Equal(randomBody, requestBody)
			return matchingResult
		})

		//
		// Set always matching static regular expression which should be ignored as there is a Matching Callback registered
		//

		mockRequest.SetBody(".*")

		//
		// Matching Callback criteria should be used ignoring always matching static regular expression
		//

		ass.Equal(matchingResult, mockRequest.BodyMatch(randomBody))
	}
}

//
// TestRequestsBodyMissingCriteriaMatching() will ensure Requests Body missing criteria is working properly
//

func TestRequestsBodyMissingCriteriaMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Create Request
	//

	mockRequest := new(go_mock_yourself_http.Request)

	//
	// Ensure Request matches any Body on missing Body matching criteria
	//

	ass.True(mockRequest.BodyMatch("missing body criteria should make any request body match"))
}
