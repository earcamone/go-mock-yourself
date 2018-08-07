package end_to_end_test

import (
	"testing"

	"github.com/mercadolibre/go-mock-yourself/http/tests/internal/e2e_helpers"
)

//
// TestURLCallbackMatchingMock() will ensure the URL Matching Callback scheme is working properly
//

func TestMethodCallbackMatchingMock(t *testing.T) {
	for _, method := range e2e_helpers.SupportedHTTPMethods {
		for _, matchingCallbackResult := range []bool{ true, false } {
			matchingCallback := func(string) bool {
				return matchingCallbackResult
			}

			//
			// Generate URL Matching Mock
			//

			mock := e2e_helpers.CreateMock(matchingCallback, nil, nil, nil)

			//
			// Generate HTTP request matching just generated mock
			//

			request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", nil, "")

			//
			// Ensure Mock matches HTTP Request
			//

			switch matchingCallbackResult {
			case true:
				mockShouldMatch(t, request, mock)

			case false:
				mockShouldNotMatch(t, request, mock)
			}
		}
	}
}

//
// TestURLCallbackMatchingMock() will ensure the URL Matching Callback scheme is working properly
//

func TestURLCallbackMatchingMock(t *testing.T) {
	for _, method := range e2e_helpers.SupportedHTTPMethods {
		for _, matchingCallbackResult := range []bool{ true, false } {
			matchingCallback := func(string) bool {
				return matchingCallbackResult
			}

			//
			// Generate URL Matching Mock
			//

			mock := e2e_helpers.CreateMock(nil, matchingCallback, nil, nil)

			//
			// Generate HTTP request matching just generated mock
			//

			request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", nil, "")

			//
			// Ensure Mock matches HTTP Request
			//

			switch matchingCallbackResult {
			case true:
				mockShouldMatch(t, request, mock)

			case false:
				mockShouldNotMatch(t, request, mock)
			}
		}
	}
}

//
// TestBodyCallbackMatchingMock() will ensure the URL Matching Callback scheme is working properly
//

func TestBodyCallbackMatchingMock(t *testing.T) {
	for _, method := range e2e_helpers.SupportedHTTPMethods {
		for _, matchingCallbackResult := range []bool{ true, false } {
			matchingCallback := func(string) bool {
				return matchingCallbackResult
			}

			//
			// Generate URL Matching Mock
			//

			mock := e2e_helpers.CreateMock(nil, nil, nil, matchingCallback)

			//
			// Generate HTTP request matching just generated mock
			//

			request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", nil, "")

			//
			// Ensure Mock matches HTTP Request
			//

			switch matchingCallbackResult {
			case true:
				mockShouldMatch(t, request, mock)

			case false:
				mockShouldNotMatch(t, request, mock)
			}
		}
	}
}

//
// TestBodyCallbackMatchingMock() will ensure the URL Matching Callback scheme is working properly
//

func TestHeadersCallbackMatchingMock(t *testing.T) {
	for _, method := range e2e_helpers.SupportedHTTPMethods {
		for _, matchingCallbackResult := range []bool{ true, false } {
			matchingCallback := func(map[string]string) bool {
				return matchingCallbackResult
			}

			//
			// Generate URL Matching Mock
			//

			mock := e2e_helpers.CreateMock(nil, nil, matchingCallback, nil)

			//
			// Generate HTTP request matching just generated mock
			//

			headers := make(map[string]string)
			headers["No-Matter-What-We-Send"] = "it will match or not depending on matchingCallback return value"

			request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", headers, "")

			//
			// Ensure Mock matches HTTP Request
			//

			switch matchingCallbackResult {
			case true:
				mockShouldMatch(t, request, mock)

			case false:
				mockShouldNotMatch(t, request, mock)
			}
		}
	}
}