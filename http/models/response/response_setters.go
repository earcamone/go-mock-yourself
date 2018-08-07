package go_mock_yourself_response

//
// ===================== Target Response Setter methods =====================
//

import (
	"fmt"
	"net/http"
)

//
// SetError() will set the Request Mock Response static error or Error Generator function
//

func (self *Response) SetError(responseError interface{}) error {
	switch responseError.(type) {
	case func(string, http.Request) error:
		//
		// Mock Response should have an Error Generator function registered?
		//

		errorGenerator := responseError.(func(string, http.Request) error)
		self.errorGenerator = &errorGenerator

	case error:
		//
		// Mock Response should have a static error registered?
		//

		self.eRROR = responseError.(error)

	default:
		return fmt.Errorf("unsupported error type. SetError() can receive either a static error or an Error Generator Callback with the signature 'func(string, http.Request) error', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}

//
// SetStatusCode() will set the Request Mock Response static body or Body Generator function
//

func (self *Response) SetStatusCode(responseStatusCode interface{}) error {
	switch responseStatusCode.(type) {
	case func(string, http.Request) int:
		//
		// Mock Response should have a Status Code Generator function registered?
		//

		statusCodeGenerator := responseStatusCode.(func(string, http.Request) int)
		self.statusCodeGenerator = &statusCodeGenerator

	case int:
		//
		// Mock Response should have a static Status Code registered?
		//

		self.statusCode = responseStatusCode.(int)

	default:
		return fmt.Errorf("unsupported status code type. SetStatusCode() can receive either a static integer status code or a Status Code Generator Callback with the signature 'func(string, http.Request) int', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}

//
// SetBody() will set the Request Mock Response static body or Body Generator function
//

func (self *Response) SetBody(responseBody interface{}) error {
	switch responseBody.(type) {
	case func(string, http.Request) string:
		//
		// Mock Response should have a Body Generator function registered?
		//

		bodyGenerator := responseBody.(func(string, http.Request) string)
		self.bodyGenerator = &bodyGenerator

	case string:
		//
		// Mock Response should have a static body registered?
		//

		self.body = responseBody.(string)

	default:
		return fmt.Errorf("unsupported body type. SetBody() can receive either a static body string or a Body Generator Callback with the signature 'func(string, http.Request) string', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}

//
// SetHeaders() will set the Request Mock Response static headers or Headers Generator function
//

func (self *Response) SetHeaders(responseHeaders interface{}) error {
	switch responseHeaders.(type) {
	case func(string, http.Request) map[string]string:
		//
		// Mock Response should have a Body Generator function registered?
		//

		headersGenerator := responseHeaders.(func(string, http.Request) map[string]string)
		self.headersGenerator = &headersGenerator

	case map[string]string:
		//
		// Mock Response should have a static body registered?
		//

		self.headers = responseHeaders.(map[string]string)

	default:
		return fmt.Errorf("unsupported headers type. SetHeaders() can receive either a static headers map (map[string]string) or a Headers Generator Callback with the signature 'func(string, http.Request) map[string]string', for a more detailed description kindly check Go Mock Yourself documentation")
	}

	return nil
}
