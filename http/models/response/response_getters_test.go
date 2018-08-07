package go_mock_yourself_response_test

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
// TestResponseGetError() will ensure Responses GetError() method is working properly
//

func TestResponseGetError(t *testing.T) {
	ass := assert.New(t)

	//
	// Create asserting error
	//

	responseError := fmt.Errorf("i'm a cute and loving error, my dad thinks i'm the best error in the world.")

	//
	// Create Response and set static error for later asserting
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetError(responseError)

	//
	// Get Response Error and ensure matches the previously set static error
	//

	mockResponseError := mockResponse.GetError("MOCK_NAME_LIKE_TRUE_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(responseError, mockResponseError)

	//
	// Set Response Error Generator Callback
	//

	mockResponse = new(go_mock_yourself_http.Response)
	mockResponse.SetError(func(string, http.Request) error {
		return responseError
	})

	//
	// Assert GetError() returns our Error Generator Callback Error
	//

	mockResponseError = mockResponse.GetError("MOCK_NAME_LIKE_TRUE_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(responseError, mockResponseError)
}

//
// TestResponseGetBody() will ensure Response get status code scheme is working properly
//

func TestResponseGetStatusCode(t *testing.T) {
	ass := assert.New(t)

	//
	// Create asserting error
	//

	randomStatusCode := go_mock_yourself_helpers.RandomInt(999999, 999999999)

	//
	// Create Response and set random status code for later asserting
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(randomStatusCode)

	//
	// Get Response status code and ensure matches the previously set random status code
	//

	mockResponseStatusCode := mockResponse.GetStatusCode("MOCK_NAME_LIKE_COOL_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(randomStatusCode, mockResponseStatusCode)

	//
	// Set Response Status Code Generator Callback
	//

	mockResponse = new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(func(string, http.Request) int {
		return randomStatusCode
	})

	//
	// Assert GetStatusCode() returns our Status Code Generator Callback status code
	//

	mockResponseStatusCode = mockResponse.GetStatusCode("MOCK_NAME_LIKE_COOL_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(randomStatusCode, mockResponseStatusCode)
}

//
// TestResponseGetBody() will ensure Response get body scheme is working properly
//

func TestResponseGetBody(t *testing.T) {
	ass := assert.New(t)

	//
	// Create asserting random body
	//

	randomBody := go_mock_yourself_helpers.RandomString(1000)

	//
	// Create Response and set random status code for later asserting
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetBody(randomBody)

	//
	// Get Response status code and ensure matches the previously set random status code
	//

	mockResponseBody := mockResponse.GetBody("MOCK_NAME_LIKE_COOL_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(randomBody, mockResponseBody)

	//
	// Set Response Status Code Generator Callback
	//

	mockResponse = new(go_mock_yourself_http.Response)
	mockResponse.SetBody(func(string, http.Request) string {
		return randomBody
	})

	//
	// Assert GetStatusCode() returns our Status Code Generator Callback status code
	//

	mockResponseBody = mockResponse.GetBody("MOCK_NAME_LIKE_COOL_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(randomBody, mockResponseBody)
}

//
// TestResponseHeaders() will ensure Response get headers scheme is working properly
//

func TestResponseHeaders(t *testing.T) {
	ass := assert.New(t)

	//
	// Create asserting random headers
	//

	okSometimesIMightMakeVariableNamesTooLargeJustSometimes := go_mock_yourself_helpers.RandomString

	randomHeaders := make(map[string]string)
	randomHeaders[okSometimesIMightMakeVariableNamesTooLargeJustSometimes(666)] = okSometimesIMightMakeVariableNamesTooLargeJustSometimes(666)
	randomHeaders[okSometimesIMightMakeVariableNamesTooLargeJustSometimes(666)] = okSometimesIMightMakeVariableNamesTooLargeJustSometimes(666)
	randomHeaders[okSometimesIMightMakeVariableNamesTooLargeJustSometimes(666)] = okSometimesIMightMakeVariableNamesTooLargeJustSometimes(666)

	//
	// Create Response and set random headers for later asserting
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetHeaders(randomHeaders)

	//
	// Get Response random headers and ensure matches the previously set random headers
	//

	mockResponseHeaders := mockResponse.GetHeaders("MOCK_NAME_LIKE_COOL_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(randomHeaders, mockResponseHeaders)

	//
	// Set Response Headers Generator Callback
	//

	mockResponse = new(go_mock_yourself_http.Response)
	mockResponse.SetHeaders(func(string, http.Request) map[string]string {
		return randomHeaders
	})

	//
	// Assert GetHeaders() returns our Headers Generator Callback random headers
	//

	mockResponseHeaders = mockResponse.GetHeaders("MOCK_NAME_LIKE_COOL_CONSTANTS_LOOK", new(http.Request))
	ass.Equal(randomHeaders, mockResponseHeaders)
}
