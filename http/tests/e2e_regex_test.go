package end_to_end_test

import (
	"testing"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http/helpers"

	// Go Mock Yourself e2e Tests Imports
	"github.com/mercadolibre/go-mock-yourself/http/tests/internal/e2e_helpers"
	"net/http"
)

//
// TestURLRegexMatchingMock() will ensure Mock Request URL Regex Matching scheme is working properly
//

func TestURLRegexMatchingMock(t *testing.T) {
	for _, method := range e2e_helpers.SupportedHTTPMethods {
		//
		// Generate URL Matching Mock
		//

		mock := e2e_helpers.CreateMock(nil, "/my-(.*)-url", nil, nil)

		//
		// Generate HTTP request matching just generated mock
		//

		request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", nil, "")

		//
		// Ensure Mock matches HTTP Request
		//

		mockShouldMatch(t, request, mock)
	}
}

//
// TestBodyRegexMatchingMock() will ensure Mock Request Body Regex Matching scheme is working properly
//

func TestBodyRegexMatchingMock(t *testing.T) {
	for _, method := range e2e_helpers.SupportedHTTPMethods {
		if method == http.MethodHead {
			continue
		}

		//
		// Generate Body Matching Mock
		//

		mock := e2e_helpers.CreateMock(nil, nil, nil, "my (.*) body")

		//
		// Generate HTTP request matching just generated mock
		//

		request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", nil, "my bogus body")

		//
		// Ensure Mock matches HTTP Request
		//

		mockShouldMatch(t, request, mock)
	}
}

//
// TestHeadersRegexMatchingMock() will ensure Mock Request Headers Regex Matching scheme is working properly
//

func TestHeadersRegexMatchingMock(t *testing.T) {
	for _, method := range e2e_helpers.SupportedHTTPMethods {
		randomString := go_mock_yourself_helpers.RandomString

		//
		// Generate random HTTP Headers
		//

		headers := make(map[string]string)
		headers[randomString(100)] = randomString(100)
		headers[randomString(100)] = randomString(100)

		//
		// Generate URL Matching Mock
		//

		mock := e2e_helpers.CreateMock(nil, nil, headers, nil)

		//
		// Generate HTTP request matching just generated mock
		//

		request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", headers, "")

		//
		// Ensure Mock matches HTTP Request
		//

		mockShouldMatch(t, request, mock)
	}
}

//
// TestHeadersRegexMatchingMock() will ensure Mock Request Headers Regex Matching scheme is working properly
//

func TestMethodsRegexMatchingMock(t *testing.T) {
	for _, mockMethod := range e2e_helpers.SupportedHTTPMethods {
		for _, requestMethod := range e2e_helpers.SupportedHTTPMethods {
			//
			// Generate URL Matching Mock
			//

			mock := e2e_helpers.CreateMock(mockMethod, nil, nil, nil)

			//
			// Generate HTTP request matching just generated mock
			//

			request := e2e_helpers.CreateHTTPRequest(requestMethod, "/my-bogus-url", nil, "")

			//
			// Ensure Mock matches HTTP Request
			//

			if requestMethod == mockMethod {
				mockShouldMatch(t, request, mock)

			} else {
				mockShouldNotMatch(t, request, mock)
			}
		}
	}
}
