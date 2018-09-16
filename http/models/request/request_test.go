package go_mock_yourself_request

import (
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"
)

//
// TestReadyRequests() will ensure Mock Requests Ready() method is working correctly
//

func TestReadyResponses(t *testing.T) {
	//
	// If Mock Request has at least one matching criteria, its ready :)
	//
	// NOTE: each specific matching criteria setter function is tested both in setters/getters tests files in this directory
	//

	ass := assert.New(t)

	//
	// Ensure Mock Request becomes ready after setting an HTTP Request Criteria Function
	//

	mockRequest := new(Request)
	ass.NotNil(mockRequest.Ready())

	mockRequest.SetRequestCriteria(func(*http.Request) bool {
		return true
	})

	ass.Nil(mockRequest.Ready())

	//
	// Ensure Mock Request becomes ready after setting an URL matching criteria
	//

	mockRequest = new(Request)
	ass.NotNil(mockRequest.Ready())

	mockRequest.SetUrl(".*")
	ass.Nil(mockRequest.Ready())

	//
	// Ensure Mock Request becomes ready after setting a Body matching criteria
	//

	mockRequest = new(Request)
	ass.NotNil(mockRequest.Ready())

	mockRequest.SetBody(".*")
	ass.Nil(mockRequest.Ready())

	//
	// Ensure Mock Request becomes ready after setting a Headers matching criteria
	//

	mockRequest = new(Request)
	ass.NotNil(mockRequest.Ready())

	matchingHeaders := make(map[string]string)
	matchingHeaders[".*"] = ".*"

	mockRequest.SetHeaders(matchingHeaders)
	ass.Nil(mockRequest.Ready())

	//
	// Ensure Mock Request becomes ready after setting a Method matching criteria
	//

	mockRequest = new(Request)
	ass.NotNil(mockRequest.Ready())

	mockRequest.SetMethod(".*")
	ass.Nil(mockRequest.Ready())
}
