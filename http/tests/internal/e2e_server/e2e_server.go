package e2e_server

import (
	"time"
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http/helpers"
	"github.com/earcamone/go-mock-yourself/http/tests/internal/e2e_helpers"
)

//
// Run() will initialize and run the e2e Testing HTTP Server
//

func Run() {
	//
	// Initialize Go Mock Yourself e2e Testing Server
	//

	Router := gin.Default()

	for _, method := range e2e_helpers.SupportedHTTPMethods {
		Router.Handle(method, "/*url", e2eTestingServerHandler(method))
	}

	//
	// Start Testing Server
	//

	go Router.Run()

	//
	// Let's give Gin enough time to run
	//

	time.Sleep(time.Second * 2)
}

func GetResponse() *Response {
	e2eServerResponseMutex.Lock()
	response := e2eServerResponse
	e2eServerResponseMutex.Unlock()

	return response
}

func SaveResponse(response *Response) {
	e2eServerResponseMutex.Lock()
	e2eServerResponse = response
	e2eServerResponseMutex.Unlock()
}

//
// TestingServerResponseMatch() will ensure the received HTTP Response matches last e2e Testing HTTP Server Response
//

func ResponseMatch(t *testing.T, httpResponse *http.Response) {
	ass := assert.New(t)

	//
	// Retrieve last Testing HTTP Server Response
	//

	e2eResponse := GetResponse()
	ass.NotNil(e2eResponse, "e2e Testing Server Response is nil! probable concurrency issue!")

	//
	// Retrieve HTTP Response information
	//

	code, headers, body := go_mock_yourself_helpers.ResponseInformation(httpResponse)

	//
	// Ensure HTTP Response matches e2e Testing HTTP Server
	//

	ass.Equal(e2eResponse.Code, code, "Request's response code doesn't match e2e Testing HTTP Server response code. Method: %s, e2e code: %d, response code: %d", httpResponse.Request.Method, e2eResponse.Code, code)

	if httpResponse.Request.Method != http.MethodHead {
		ass.Equal(e2eResponse.Body, body, "Request's response body doesn't match e2e Testing HTTP Server response body. Method: %s, e2e body: %s, response body: %s", httpResponse.Request.Method, e2eResponse.Body, body)
	}

	ass.True(headersContainHeaders(headers, e2eResponse.Headers),"Request's response headers don't match e2e Testing HTTP Server response headers. Method: %s, e2e headers: %v, response headers: %v", httpResponse.Request.Method, e2eResponse.Headers, headers)
}

func headersContainHeaders(headers map[string]string, expectedHeaders map[string]string) bool {
	if len(expectedHeaders) == 0 {
		return true
	}

	for expectedHeader, expectedHeaderValue := range expectedHeaders {
		expectedHeaderIsPresent := false

		for header, value := range headers {
			if header == expectedHeader && value == expectedHeaderValue {
				expectedHeaderIsPresent = true
				break
			}
		}

		if expectedHeaderIsPresent == false {
			return false
		}
	}

	return true
}