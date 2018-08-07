package go_mock_yourself_response

//
// ===================== Target Response Getter methods =====================
//

import (
	"net/http"
)

//
// GetError() will return the Mock Response error based on the received Mock name and HTTP request
//

func (self *Response) GetError(mockName string, request *http.Request) error {
	errorResponse := self.eRROR

	if self.errorGenerator != nil {
		errorResponse = (*self.errorGenerator)(mockName, *request)
	}

	return errorResponse
}

//
// GetStatusCode() will return the Mock Response status code
//

func (self *Response) GetStatusCode(mockName string, request *http.Request) int {
	responseStatusCode := self.statusCode

	if self.statusCodeGenerator != nil {
		responseStatusCode = (*self.statusCodeGenerator)(mockName, *request)
	}

	return responseStatusCode
}

//
// GetBody() will return the Mock Response body
//

func (self *Response) GetBody(mockName string, request *http.Request) string {
	responseBody := self.body

	if self.bodyGenerator != nil {
		responseBody = (*self.bodyGenerator)(mockName, *request)
	}

	return responseBody
}

//
// GetBody() will return the Mock Response body
//

func (self *Response) GetHeaders(mockName string, request *http.Request) map[string]string {
	responseHeaders := self.headers

	if self.headersGenerator != nil {
		responseHeaders = (*self.headersGenerator)(mockName, *request)
	}

	return responseHeaders
}
