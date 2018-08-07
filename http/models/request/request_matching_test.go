package go_mock_yourself_request

import (
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"
)

//
// TestMatchingScheme() will ensure Mock Requests main matching function Match() is working correctly
//

func TestMatchingScheme(t *testing.T) {
	ass := assert.New(t)

	//
	// Create matching criterias dynamic failures slice (each criteria will return the bool in a specific slice index)
	//

	dynamicMatchingResults := []bool{ true, true, true, true }

	//
	// Build Mock Request returning bool slice values
	//

	mockRequest := new(Request)

	mockRequest.SetMethod(func(method string) bool {
		return dynamicMatchingResults[0]
	})

	mockRequest.SetUrl(func(url string) bool {
		return dynamicMatchingResults[1]
	})

	mockRequest.SetHeaders(func(headers map[string]string) bool {
		return dynamicMatchingResults[2]
	})

	mockRequest.SetBody(func(body string) bool {
		return dynamicMatchingResults[3]
	})

	//
	// Right now all registered matching criterias are returning true so Mock Request should match
	//

	request := new(http.Request)
	ass.True(mockRequest.Match(request))

	//
	// Let's now ensure Match() is failing properly
	//

	for boolSliceIndex := range dynamicMatchingResults {
		//
		// Update dynamicMatchingResults making ONLY ONE matching criteria return false
		//

		for i := 0; i < len(dynamicMatchingResults); i++ {
			dynamicMatchingResults[i] = boolSliceIndex != i
		}

		//
		// When at least one matching criteria returns false, Mock Request should not match
		//

		ass.False(mockRequest.Match(request))
	}

	//
	// As we are registering an HTTP Request Dynamic Matching Function only its result
	// should determine if the Mock Request is matching or not the incoming request
	//

	for _, matchingRequestCriteria := range []bool{ true, false } {
		mockRequest.SetRequestCriteria(func(*http.Request) bool {
			return matchingRequestCriteria
		})

		ass.Equal(matchingRequestCriteria, mockRequest.Match(request))
	}
}
