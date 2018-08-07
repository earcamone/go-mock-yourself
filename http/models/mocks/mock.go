package go_mock_yourself_mocks

import (
	"fmt"
	"time"
	"bufio"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http/models/request"
	"github.com/mercadolibre/go-mock-yourself/http/models/response"
)

//
//
//

type Mock struct {
	//
	// Mock name
	//

	Name string

	//
	// Mock stream should be logged?
	//

	Logging *bufio.Writer

	//
	// Request processing time-out?
	//

	Timeout time.Duration

	//
	// Mock Request/Target Information
	//

	Request *go_mock_yourself_request.Request

	//
	// Mock Response Information
	//

	Response *go_mock_yourself_response.Response
}

func New(name string, request *go_mock_yourself_request.Request, response *go_mock_yourself_response.Response, optionalLogging ...*bufio.Writer) (*Mock, error) {
	//
	// Ensure Mock name is correct
	//

	if len(name) == 0 {
		return nil, fmt.Errorf("Invalid Mock Name. Kindly specify an identification name for the mock, its name is sent to all Generator Callbacks together with the target HTTP Request so you can conditionally build your response based on the Mock name also")
	}

	//
	// Ensure Mock Request is correct
	//

	if request == nil {
		return nil, fmt.Errorf("Mock Request can't be nil, you would be installing a Mock which would never match anything")
	}

	if requestError := request.Ready(); requestError != nil {
		return nil, fmt.Errorf("Invalid Mock Request. You are probably initializing it incorrectly, kindly check documentation: %s", requestError.Error())
	}

	//
	// Build mock
	//

	mock := Mock {
		Name: name,
		Request: request,
		Response: response,
	}

	//
	// Optional logging stream specified?
	//

	if len(optionalLogging) > 0 {
		mock.Logging = optionalLogging[0]
	}

	return &mock, nil
}

func (self Mock) Ready() error {
	if self.Request == nil {
		return fmt.Errorf("Mock not initialized properly, it requires at least a Mock Matching Request")
	}

	if requestError := self.Request.Ready(); requestError != nil {
		return requestError
	}

	if self.Response != nil {
		if responseError := self.Response.Ready(); responseError != nil {
			return responseError
		}
	}

	return nil
}
