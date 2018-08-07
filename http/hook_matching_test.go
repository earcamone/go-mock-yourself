package go_mock_yourself_http

import (
	"fmt"
	"bufio"
	"bytes"
	"testing"
	"strings"
	"net/http"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
)

//
// TestHookPauseScheme() ensures Go Mock Yourself Pausing Scheme is working correctly
//

func TestHookPauseScheme(t *testing.T) {
	ass := assert.New(t)

	//
	// Reset Mocking Scheme Mocks
	//

	Reset()

	//
	// Install always matching Mock
	//

	matchingMock := createMatchingMock()
	Install(matchingMock)

	//
	// Create bogus HTTP Request to fed getMatchingMock()
	//

	httpRequest := createHTTPRequest(http.MethodGet, "http://notengoenie.com/", nil, "")

	//
	// Fed getMatchingMock() and ensure it returns the installed mock
	//

	ass.Equal(matchingMock, getMatchingMock(httpRequest))

	//
	// Pause Mocking Scheme!
	//

	Pause()

	//
	// Call again getMatchingMock() and this time ensure it DOES NOT return the matching mock
	//

	ass.Nil(getMatchingMock(httpRequest))

	//
	// Play Mocking Scheme!
	//

	Play()

	//
	// Call again getMatchingMock() and this time ensure it DOES NOT return the matching mock
	//

	ass.Equal(matchingMock, getMatchingMock(httpRequest))
}

//
// TestMatchingMockScheme() will ensure getMatchingMock() function is working properly
//

func TestMatchingMockScheme(t *testing.T) {
	ass := assert.New(t)

	//
	// Reset Mocking Scheme Mocks
	//

	Reset()

	//
	// Create bogus HTTP Request to fed getMatchingMock()
	//

	randomString := go_mock_yourself_helpers.RandomString(100)

	expectedHeaders := make(map[string]string)
	expectedHeaders[randomString] = randomString

	expectedBody := fmt.Sprintf("Rasti loves mindcraft, morse code, olorin and off course random data in tests: %s", randomString)
	httpRequest := createHTTPRequest(http.MethodGet, "http://notengoenie.com/", expectedHeaders, expectedBody)

	//
	// Build Mock Request with dynamic matching criterias using Dynamic Matching Callbacks, additionally ensure
	// inside each Dynamic Matching Callback the corresponding parameters are received based on httpRequest values
	//

	dynamicMatchingResults := []bool{ true, true, true, true }
	mockRequest := new(Request)

	mockRequest.SetMethod(func(method string) bool {
		ass.Equal(httpRequest.Method, method)
		return dynamicMatchingResults[0]
	})

	mockRequest.SetUrl(func(url string) bool {
		ass.Equal(httpRequest.URL.String(), url)
		return dynamicMatchingResults[1]
	})

	mockRequest.SetHeaders(func(headers map[string]string) bool {
		//
		// NOTE: we are checking expectedHeaders are included in headers and not just ensuring they are equal as
		// internally Go Mock Yourself HTTP requests handler mimik Go native's package handler which if some specific
		// headers are missing in the Request Transport, they are transparently added (for example Host and Client).
		//

		ass.True(HeadersIncludeHeaders(headers, expectedHeaders))
		return dynamicMatchingResults[2]
	})

	mockRequest.SetBody(func(body string) bool {
		ass.Equal(expectedBody, body)
		return dynamicMatchingResults[3]
	})

	//
	// If no Mocks are installed, getMatchingMock() should never return a mock
	//

	ass.Nil(getMatchingMock(httpRequest))

	//
	// Install Mock with Dynamic Matching Request (currently all dynamicMatchingResults are true)
	//
	// NOTE: You should always check errors against Install(), but you are you and me is me, mmm, it doesn't work like this right?
	//

	dynamicMatchingMock, _ := NewMock("dynamic matching mock", mockRequest, nil)
	Install(dynamicMatchingMock)

	//
	// getMatchingMock() should match as all matching criterias will return true according to current dynamicMatchingResults values
	//

	ass.Equal(dynamicMatchingMock, getMatchingMock(httpRequest))

	//
	// Let's now make each matching criteria return false individually while all others return true,
	// getMatchingMock() should then return nil as Mocks should match only if ALL matching criterias match
	//

	for boolSliceIndex := range dynamicMatchingResults {
		//
		// Update dynamicMatchingResults making ONLY ONE matching criteria return false
		//

		for i := 0; i < len(dynamicMatchingResults); i++ {
			dynamicMatchingResults[i] = boolSliceIndex != i
		}

		//
		// When any matching criteria does not match, getMatchingMock() should not return Mock
		//

		ass.Nil(getMatchingMock(httpRequest))
	}

	//
	// Let's be a bad anarchist
	//

	Reset()
}

//
// createMatchingMock() will create an always matching mock
//

func createMatchingMock(optionalShouldMock ...func(*http.Request) bool) *Mock {
	mockRequest := new(Request)
	mockRequest.SetUrl(".*")

	if len(optionalShouldMock) > 0 {
		mockRequest.SetShouldMock(optionalShouldMock[0])
	}

	mockResponse := new(Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody("bichi has a good heart, but he is usually infumable (you should pronounce this word in english, try it)")

	loggingBuffer := bufio.NewWriter(bytes.NewBuffer([]byte("")))
	mock, _ := NewMock("always matching mock", mockRequest, mockResponse, loggingBuffer)

	return mock
}

//
// createHTTPRequest() will build an http.Request from the received parameters
//

func createHTTPRequest(method string, url string, headers map[string]string, body string) *http.Request {
	request, _ := http.NewRequest(method, url, strings.NewReader(body))

	if len(headers) > 0 {
		request.Header = go_mock_yourself_helpers.HeadersMapToHTTPHeaders(headers)
	}

	return request
}

//
// HeadersIncludeHeaders() will return true if the received matchingHeaders map is included in the received allHeaders map
//

func HeadersIncludeHeaders(allHeaders map[string]string, expectedHeaders map[string]string) bool {
	for expectedHeader, expectedValue := range expectedHeaders {
		matchingHeader := false

		for header, value := range allHeaders {
			if header == expectedHeader && value == expectedValue {
				matchingHeader = true
				break
			}
		}

		if matchingHeader == false {
			return false
		}
	}

	return true
}
