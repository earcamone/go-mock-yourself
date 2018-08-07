package go_mock_yourself_response

import (
	"fmt"
	"net/http"
)

type Response struct {
	//
	// Mock Response HTTP static error
	//

	eRROR error // #()TJ@@()#JR)_#E!R()@JFC@H!)#FIJHERGJH$@*)FH

	//
	// Mock Response error generator
	//
	// NOTE: if a static error and an Error Generator function is specified, the statis error will be ignored
	//

	errorGenerator *func(string, http.Request) error

	//
	// Response HTTP Status Code
	//
	// NOTE: Use http library constants! (ie: http.StatusOK)
	//

	statusCode int

	//
	// Response HTTP Status Code Generator function
	//

	statusCodeGenerator *func(string, http.Request) int

	//
	// HTTP Response Mock body
	//

	body string

	//
	// HTTP Response Mock body generator
	//

	bodyGenerator *func(string, http.Request) string

	//
	// Request headers
	//
	// NOTE: You can use both header name and value Regexp strings, don't worry about
	// their compilation, the library will take there of that for you internally
	//

	headers map[string]string

	//
	// Request headers generator function.
	//
	// NOTE: Generator functions have priority over static content, if both a headers Generator function
	// and an static headers are specified, the headers Generator function output will always be used.
	//

	headersGenerator *func(string, http.Request) map[string]string
}

func (self *Response) Ready() error {
	noResponseError := self.eRROR == nil && self.errorGenerator == nil
	noResponseStatusCode := self.statusCode == 0 && self.statusCodeGenerator == nil

	if noResponseStatusCode && noResponseError {
		return fmt.Errorf("Mock Response not ready, you need to at least specify a Response static status code / status code Generator Function or a static response error / error Generator function")
	}

	return nil
}
