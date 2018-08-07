package go_mock_yourself_request_test

import (
	"fmt"
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http"
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
)

//
// Supported HTTP Methods
//

var httpMethods = []string {
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

//
// TestMethodsMatchingScheme() will ensure SetMethod() regular expressions scheme is working properly
//

func TestMethodsRegexMatchingScheme(t *testing.T) {
	ass := assert.New(t)

	//
	// Set regular expressions HTTP methods
	//

	for _, matchingMethod := range httpMethods {
		//
		// Generate random matching method
		//

		randomString := go_mock_yourself_helpers.RandomString(100)
		specificMatchingRegex := fmt.Sprintf(".* %s .*", matchingMethod)

		//
		// Set request random matching method
		//

		mockRequest := new(go_mock_yourself_http.Request)
		mockRequest.SetMethod(specificMatchingRegex)

		//
		// Ensure regular expression method matching scheme is working properly
		//

		for _, currentMethod := range httpMethods {
			randomMethod := fmt.Sprintf("%s %s %s", randomString, currentMethod, randomString)
			ass.Equal(currentMethod == matchingMethod, mockRequest.MethodMatch(randomMethod))
		}
	}
}

//
// TestMethodsCallbackMatchingScheme() will ensure Requests Methods Matching Callback scheme is working properly
//

func TestMethodsCallbackMatchingScheme(t *testing.T) {
	ass := assert.New(t)

	//
	// Generate random URL to assert inside Request Matching Callback its being called properly
	//

	randomMethod := go_mock_yourself_helpers.RandomString(100)

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

		mockRequest.SetMethod(func(requestMethod string) bool {
			ass.Equal(randomMethod, requestMethod)
			return matchingResult
		})

		//
		// Set always matching static regular expression which should be ignored as there is a Matching Callback registered
		//

		mockRequest.SetMethod(".*")

		//
		// Matching Callback criteria should be used ignoring always matching static regular expression
		//

		ass.Equal(matchingResult, mockRequest.MethodMatch(randomMethod))
	}
}

//
// TestRequestsMethodsMissingCriteriaMatching() will ensure Requests Methods scheme is working properly on missing criteria
//

func TestRequestsMethodsMissingCriteriaMatching(t *testing.T) {
	ass := assert.New(t)

	//
	// Create Request
	//

	mockRequest := new(go_mock_yourself_http.Request)

	//
	// Ensure Request matches any HTTP Method on missing URL matching criteria
	//

	for _, method := range httpMethods {
		ass.True(mockRequest.MethodMatch(method))
	}
}
