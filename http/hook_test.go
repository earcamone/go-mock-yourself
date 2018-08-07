package go_mock_yourself_http

import (
	"fmt"
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
)

//
// TestShouldMockScheme() will ensure the "Should Mock" scheme on matching mocks is working correctly
//
// NOTE: Requests with Matching Mocks with registered ShouldMock() callback returning false or no registered Mock
// Response (usually used in debugging scenarios) should be redirected directly to Go's native HTTP Do() method
//

func TestShouldMockScheme(t *testing.T) {
	ass := assert.New(t)

	//
	// goMockYourselfDo() will call createHttpResponseFromMock() when attempting to mock a response, otherwise it will
	// call Go's native HTTP package client's Do() method. As we can't hook native Do() method as its already hooked
	// we are hooking createHttpResponseFromMock to assert when goMockYourselfDo() is actually handling the HTTP
	// response or is redirecting it to Go's native Do() method (a.k.a not calling createHttpResponseFromMock)
	//

	var createHttpResponseFromMockHook *monkey.PatchGuard

	createHttpResponseFromMockCalled := false
	createHttpResponseFromMockHook = monkey.Patch(createHttpResponseFromMock, func(request *http.Request, mock Mock) (*http.Response, error) {
		createHttpResponseFromMockCalled = true

		createHttpResponseFromMockHook.Unpatch()
		response, responseError := createHttpResponseFromMock(request, mock)
		createHttpResponseFromMockHook.Restore()

		return response, responseError
	})

	//
	// Create bogus HTTP Request to fed goMockYourselfDo()
	//

	expectedBody := fmt.Sprintf("i'm a cute loving body with off course random data: %s", go_mock_yourself_helpers.RandomString(100))
	httpRequest := createHTTPRequest(http.MethodGet, "http://notengoenie.com/", nil, expectedBody)

	//
	// Install always matching mock, thus createHttpResponseFromMock() should always be called
	//

	matchingMock := createMatchingMock()
	Install(matchingMock)

	//
	// createHttpResponseFromMock() should had been called as no ShouldMock method is installed
	//

	goMockYourselfDo(http.DefaultClient, httpRequest)
	ass.True(createHttpResponseFromMockCalled)

	//
	// Reset Go Mock Yourself Mocks
	//

	Reset()

	//
	// Reset createHttpResponseFromMock() calling indicator
	//

	createHttpResponseFromMockCalled = false

	//
	// Remove always matching Mock Response, thus Go Mock Yourself should not mock response
	//

	matchingMock.Response = nil

	//
	// createHttpResponseFromMock() should had NOT been called as matching Mock does not have a Mock Response
	//

	goMockYourselfDo(http.DefaultClient, httpRequest)
	ass.False(createHttpResponseFromMockCalled)

	//
	// Install conditionally matching mock through ShouldMock scheme
	//

	for _, shouldMock := range []bool{ true, false } {
		//
		// Reset Go Mock Yourself Mocks
		//

		Reset()

		//
		// Install an always matching mock with a ShouldMock Callback returning shouldMock loop value
		//

		Install(createMatchingMock(func(*http.Request) bool {
			return shouldMock
		}))

		//
		// Reset createHttpResponseFromMock() calling indicator
		//

		createHttpResponseFromMockCalled = false

		//
		// createHttpResponseFromMock() should be called depending on ShouldMock returned value (loop bool value)
		//

		goMockYourselfDo(http.DefaultClient, httpRequest)
		ass.Equal(shouldMock, createHttpResponseFromMockCalled)
	}
}
