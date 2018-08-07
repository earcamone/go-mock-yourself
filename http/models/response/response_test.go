package go_mock_yourself_response

import (
	"fmt"
	"testing"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"
)

//
// TestReadyResponses() will ensure Responses Ready() method is working correctly
//

func TestReadyResponses(t *testing.T) {
	ass := assert.New(t)

	//
	// If Mock Response has at least a Response Code, its ready :)
	//
	// NOTE: SetError() dynamic/static error setting is tested in Getters tests (response_getters_test.go)
	//

	mockResponse := new(Response)
	ass.NotNil(mockResponse.Ready())

	mockResponse.SetError(fmt.Errorf("my lovely bogus error"))
	ass.Nil(mockResponse.Ready())

	//
	// If Mock Response is set to fail with a Go error, its ready :)
	//
	// NOTE: SetStatusCode() dynamic/static code setting is tested in Getters tests (response_getters_test.go)
	//

	mockResponse = new(Response)
	ass.NotNil(mockResponse.Ready())

	mockResponse.SetStatusCode(666)
	ass.Nil(mockResponse.Ready())

	//
	// When Mock Response has no Response Code or either set a Go error to fail, its not ready!
	//

	mockResponse = new(Response)
	ass.NotNil(mockResponse.Ready())

	bogusHeaders := make(map[string]string)

	mockResponse.SetBody("my lovely bogus body")
	mockResponse.SetHeaders(bogusHeaders)

	ass.NotNil(mockResponse.Ready())
}