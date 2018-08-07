package go_mock_yourself_response_test

//
// =========================== Response Setters Tests ==============================
//
// NOTE: Remaining Setters schemes are tested indirectly in the Getters() tests
//

import (
	"fmt"
	"testing"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http"
	"github.com/mercadolibre/go-mock-yourself/http/models/internal"
)

//
// TestResponseErrorSettingErrorHandling() will ensure Responses SetStatusCode() is handling errors properly
//

func TestResponseErrorSettingErrorHandling(t *testing.T) {
	ass := assert.New(t)

	//
	// Loop through all available incorrect parameters
	//

	for _, incorrectParameter := range go_mock_yourself_models_internal_helpers.IncorrectErrorParameters {
		//
		// Create Response and set invalid status code type for later assertion
		//

		mockResponse := new(go_mock_yourself_http.Response)
		mockResponseError := mockResponse.SetError(incorrectParameter)

		//
		// Ensure SetStatusCode() returned the corresponding error
		//

		expectedError := fmt.Errorf("unsupported error type. SetError() can receive either a static error or an Error Generator Callback with the signature 'func(string, http.Request) error', for a more detailed description kindly check Go Mock Yourself documentation")
		ass.Equal(expectedError, mockResponseError)
	}
}

//
// TestResponseStatusCodeSettingErrorHandling() will ensure Responses SetStatusCode() is handling errors properly
//

func TestResponseStatusCodeSettingErrorHandling(t *testing.T) {
	ass := assert.New(t)

	//
	// Loop through all available incorrect parameters
	//

	for _, incorrectParameter := range go_mock_yourself_models_internal_helpers.IncorrectIntegerParameters {
		//
		// Create Response and set invalid type for later assertion
		//

		mockResponse := new(go_mock_yourself_http.Response)
		mockResponseError := mockResponse.SetStatusCode(incorrectParameter)

		//
		// Ensure SetStatusCode() returned the corresponding error
		//

		expectedError := fmt.Errorf("unsupported status code type. SetStatusCode() can receive either a static integer status code or a Status Code Generator Callback with the signature 'func(string, http.Request) int', for a more detailed description kindly check Go Mock Yourself documentation")
		ass.Equal(expectedError, mockResponseError)
	}
}

//
// TestResponseBodySettingErrorHandling() will ensure Requests SetBody() is handling errors properly
//

func TestResponseBodySettingErrorHandling(t *testing.T) {
	ass := assert.New(t)

	//
	// Loop through all available incorrect parameters
	//

	for _, incorrectParameter := range go_mock_yourself_models_internal_helpers.IncorrectStringParameters {
		//
		// Create Response and set invalid body type for later assertion
		//

		mockResponse := new(go_mock_yourself_http.Response)
		mockResponseError := mockResponse.SetBody(incorrectParameter)

		//
		// Ensure SetBody() returned the corresponding error
		//

		expectedError := fmt.Errorf("unsupported body type. SetBody() can receive either a static body string or a Body Generator Callback with the signature 'func(string, http.Request) string', for a more detailed description kindly check Go Mock Yourself documentation")
		ass.Equal(expectedError, mockResponseError)
	}
}

//
// TestRequestsHeadersSettingErrorHandling() will ensure Requests SetHeaders() is handling errors properly
//

func TestResponsesHeadersSettingErrorHandling(t *testing.T) {
	ass := assert.New(t)

	//
	// Loop through all available incorrect parameters
	//

	for _, incorrectParameter := range go_mock_yourself_models_internal_helpers.IncorrectRegexParameters {
		//
		// Create Request and set invalid headers type for later assertion
		//

		mockRequest := new(go_mock_yourself_http.Request)
		mockRequestError := mockRequest.SetHeaders(incorrectParameter)

		//
		// Ensure SetHeaders() returned the corresponding error
		//

		expectedError := fmt.Errorf("unsupported headers type. SetHeaders() can receive either a static regular expressions headers map (map[string]string) or a Headers Matching Callback with the signature 'func(map[string]string) bool', for a more detailed description kindly check Go Mock Yourself documentation")
		ass.Equal(expectedError, mockRequestError)
	}
}
