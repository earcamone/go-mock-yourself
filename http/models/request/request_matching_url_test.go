package go_mock_yourself_request_test

import (
	"fmt"
	"testing"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http"
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
)

//
// TestRequestsURLMatching() will ensure Requests URLs regular expressions matching is working properly
//

func TestRequestsURLRegexMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Set always matching URL regular expression
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl("http://.*/")

	for i := 0; i < 100; i++ {
		randomString := go_mock_yourself_helpers.RandomString(10)
		randomUrl := fmt.Sprintf("http://%s.com/", randomString)

		ass.True(mockRequest.URLMatch(randomUrl))
	}

	//
	// Set never matching URL regular expression
	//

	mockRequest.SetUrl("http://thisurlwillnevermatch/")

	for i := 0; i < 100; i++ {
		randomString := go_mock_yourself_helpers.RandomString(100)
		randomUrl := fmt.Sprintf("http://%s.com/", randomString)

		ass.False(mockRequest.URLMatch(randomUrl))
	}

	//
	// Set specific URLs regular expressions matching
	//

	mockRequest.SetUrl("http://.*google.com/")
	ass.True(mockRequest.URLMatch("http://www.google.com/"))
}

//
// TestRequestsURLCallbackMatching() will ensure Requests URLs Matching Callback scheme is working properly
//

func TestRequestsURLCallbackMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Generate random URL to assert inside Request Matching Callback its being called properly
	//

	randomString := go_mock_yourself_helpers.RandomString(100)
	randomUrl := fmt.Sprintf("http://%s.com/", randomString)

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

		mockRequest.SetUrl(func(requestUrl string) bool {
			ass.Equal(randomUrl, requestUrl)
			return matchingResult
		})

		//
		// Set always matching static regular expression which should be ignored as there is a Matching Callback registered
		//

		mockRequest.SetUrl(".*")

		//
		// Matching Callback criteria should be used ignoring always matching static regular expression
		//

		ass.Equal(matchingResult, mockRequest.URLMatch(randomUrl))
	}
}

//
// TestRequestsURLCallbackMatching() will ensure Requests URLs Matching Callback scheme is working properly
//

func TestRequestsURLMissingCriteriaMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Create Request
	//

	mockRequest := new(go_mock_yourself_http.Request)

	//
	// Ensure Request matches any URL on missing URL matching criteria
	//

	ass.True(mockRequest.URLMatch("http://www.google.com/"))
}
