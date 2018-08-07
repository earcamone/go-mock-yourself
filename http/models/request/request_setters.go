package go_mock_yourself_request

//
// ===================== Target Requests Setter Methods =====================
//

import (
	"fmt"
	"regexp"
	"net/http"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http/models/internal"
)


//
// SetRequestCriteria() is used to set incoming HTTP Requests Matching function over the incoming http.Request object itself
//

func (self *Request) SetRequestCriteria(requestCriteria func(*http.Request) bool) error {
	self.matchingRequest = &requestCriteria
	return nil
}

//
// SetMethod() is used to set Request Target Methods matching criteria
//

func (self *Request) SetMethod(methodsMatching interface{}) error {
	switch methodsMatching.(type) {
	case func(string) bool:
		matchingFunction := methodsMatching.(func(string) bool)
		self.matchingMethods = &matchingFunction

	case string:
		method := methodsMatching.(string)

		//
		// Compile methods regular expression
		//

		compiledRegexp, regexpError := regexp.Compile(method)

		if regexpError != nil {
			return fmt.Errorf("Mock Request regular expression seems to be invalid: %s. Mock can't be installed until this regular expression is fixed", regexpError.Error())
		}

		self.methods = compiledRegexp

	default:
		return fmt.Errorf("unsupported method type. SetMethod() can receive either a regular expression string method or a Method Matching Callback with the signature 'func(string) bool', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}

//
// SetUrl() is used to set Request Target URL matching criteria
//

func (self *Request) SetUrl(urlMatching interface{}) error {
	switch urlMatching.(type) {
	case func(string) bool:
		matchingFunction := urlMatching.(func(string) bool)
		self.matchingUrl = &matchingFunction

	case string:
		url := urlMatching.(string)

		//
		// Compile url regular expression
		//

		compiledRegexp, regexpError := regexp.Compile(url)

		if regexpError != nil {
			return fmt.Errorf("Mock Request regular expression seems to be invalid: %s. Mock can't be installed until this regular expression is fixed", regexpError.Error())
		}

		self.url = compiledRegexp

	default:
		return fmt.Errorf("unsupported url type. SetUrl() can receive either a regular expression string url or an URL Matching Callback with the signature 'func(string) bool', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}

//
// SetBody() is used to set Request Target Body matching criteria
//

func (self *Request) SetBody(bodyMatching interface{}) error {
	switch bodyMatching.(type) {
	case func(string) bool:
		matchingFunction := bodyMatching.(func(string) bool)
		self.matchingBody = &matchingFunction

	case string:
		body := bodyMatching.(string)

		//
		// Compile body regular expression
		//

		compiledRegexp, regexpError := regexp.Compile(body)

		if regexpError != nil {
			return fmt.Errorf("Mock Request regular expression seems to be invalid: %s. Mock can't be installed until this regular expression is fixed", regexpError.Error())
		}

		self.body = compiledRegexp

	default:
		return fmt.Errorf("unsupported body type. SetBody() can receive either a regular expression string body or a Body Matching Callback with the signature 'func(string) bool', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}

//
// SetHeaders() is used to set Request Target Headers matching criteria
//

func (self *Request) SetHeaders(headersMatching interface{}) error {
	switch headersMatching.(type) {
	case func(map[string]string) bool:
		matchingFunction := headersMatching.(func(map[string]string) bool)
		self.matchingHeaders = &matchingFunction

	case map[string]string:
		headers := headersMatching.(map[string]string)

		//
		// Compile Headers regexes
		//

		compiledHeaders, compileError := go_mock_yourself_models_internal_helpers.CompileHeaders(headers)

		if compileError != nil {
			return fmt.Errorf("there was an error attempting to compile the received headers: %s", compileError.Error())
		}

		self.headers = compiledHeaders

	default:
		return fmt.Errorf("unsupported headers type. SetHeaders() can receive either a static regular expressions headers map (map[string]string) or a Headers Matching Callback with the signature 'func(map[string]string) bool', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}

func (self *Request) SetShouldMock(shouldMockCallback func(*http.Request) bool) {
	self.shouldMock = shouldMockCallback
}
