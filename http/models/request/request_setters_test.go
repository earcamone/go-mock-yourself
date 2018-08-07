package go_mock_yourself_request_test

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
// TestRequestsMethodsSettingErrorHandling() will ensure Requests SetMethod() is handling errors properly
//

func TestRequestsMethodsSettingErrorHandling(t *testing.T) {
	ass := assert.New(t)

	//
	// Loop through all available incorrect SetMethod() parameters
	//

	for _, incorrectParameter := range go_mock_yourself_models_internal_helpers.IncorrectRegexParameters {
		//
		// Create Request and set invalid method type for later assertion
		//

		mockRequest := new(go_mock_yourself_http.Request)
		mockRequestError := mockRequest.SetMethod(incorrectParameter)

		//
		// Ensure SetMethod() returned the corresponding error
		//

		var expectedError error

		if incorrectParameter != "(this is an invalid regex" {
			expectedError = fmt.Errorf("unsupported method type. SetMethod() can receive either a regular expression string method or a Method Matching Callback with the signature 'func(string) bool', for a more detailed description kindly check Go Mock Yourself documentation")

		} else {
			expectedError = fmt.Errorf("Mock Request regular expression seems to be invalid: error parsing regexp: missing closing ): `(this is an invalid regex`. Mock can't be installed until this regular expression is fixed")
		}

		ass.Equal(expectedError, mockRequestError)
	}
}

//
// TestRequestsURLSettingErrorHandling() will ensure Requests SetUrl() is handling errors properly
//

func TestRequestsURLSettingErrorHandling(t *testing.T) {
	ass := assert.New(t)

	//
	// Loop through all available incorrect SetUrl() parameters
	//

	for _, incorrectParameter := range go_mock_yourself_models_internal_helpers.IncorrectRegexParameters {
		//
		// Create Request and set invalid url type for later assertion
		//

		mockRequest := new(go_mock_yourself_http.Request)
		mockRequestError := mockRequest.SetUrl(incorrectParameter)

		//
		// Ensure SetUrl() returned the corresponding error
		//

		var expectedError error

		if incorrectParameter != "(this is an invalid regex" {
			expectedError = fmt.Errorf("unsupported url type. SetUrl() can receive either a regular expression string url or an URL Matching Callback with the signature 'func(string) bool', for a more detailed description kindly check Go Mock Yourself documentation")

		} else {
			expectedError = fmt.Errorf("Mock Request regular expression seems to be invalid: error parsing regexp: missing closing ): `(this is an invalid regex`. Mock can't be installed until this regular expression is fixed")
		}

		ass.Equal(expectedError, mockRequestError)
	}
}

//
// TestRequestsBodySettingErrorHandling() will ensure Requests SetBody() is handling errors properly
//

func TestRequestsBodySettingErrorHandling(t *testing.T) {
	ass := assert.New(t)

	//
	// Loop through all available incorrect SetBody() parameters
	//

	for _, incorrectParameter := range go_mock_yourself_models_internal_helpers.IncorrectRegexParameters {
		//
		// Create Request and set invalid body type for later assertion
		//

		mockRequest := new(go_mock_yourself_http.Request)
		mockRequestError := mockRequest.SetBody(incorrectParameter)

		//
		// Ensure SetBody() returned the corresponding error
		//

		var expectedError error

		if incorrectParameter != "(this is an invalid regex" {
			expectedError = fmt.Errorf("unsupported body type. SetBody() can receive either a regular expression string body or a Body Matching Callback with the signature 'func(string) bool', for a more detailed description kindly check Go Mock Yourself documentation")

		} else {
			expectedError = fmt.Errorf("Mock Request regular expression seems to be invalid: error parsing regexp: missing closing ): `(this is an invalid regex`. Mock can't be installed until this regular expression is fixed")
		}

		ass.Equal(expectedError, mockRequestError)
	}
}

//
// TestRequestsHeadersSettingErrorHandling() will ensure Requests SetHeaders() is handling errors properly
//

func TestRequestsHeadersSettingErrorHandling(t *testing.T) {
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
