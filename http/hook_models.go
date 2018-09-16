package go_mock_yourself_http

import (
	"bufio"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http/models/mocks"
	"github.com/earcamone/go-mock-yourself/http/models/request"
	"github.com/earcamone/go-mock-yourself/http/models/response"
)

//
// Models Type Aliases to make Go Mock Yourself API cooler :P
//

type Mock = go_mock_yourself_mocks.Mock
type Mocks = go_mock_yourself_mocks.Mocks

type Request = go_mock_yourself_request.Request
type Response = go_mock_yourself_response.Response

//
// Mock Builder
//

func NewMock(name string, request *Request, response *Response, optionalLogging ...*bufio.Writer) (*Mock, error) {
	return go_mock_yourself_mocks.New(name, request, response, optionalLogging...)
}

//
// NewDynamicResponse() will build a Dynamic HTTP Mock Response that will return the received Mock Responses in the
// order they were received. When the last response content is returned, if subsequent incoming requests match the Mock,
// the returned responses content loop will restart.
//

func NewDynamicResponse(responses ...*Response) *Response{
	return go_mock_yourself_response.DynamicResponse(responses...)
}
