package e2e_helpers

import (
	"github.com/mercadolibre/go-mock-yourself/http"
)

//
// CreateMock() will create a random response Mock with the received request information
//

func CreateMock(method interface{}, url interface{}, headers interface{}, body interface{}) *go_mock_yourself_http.Mock {
	mockRequest := new(go_mock_yourself_http.Request)

	if method != nil {
		mockRequest.SetMethod(method)
	}

	if url != nil {
		mockRequest.SetUrl(url)
	}

	if body != nil {
		mockRequest.SetBody(body)
	}

	if headers != nil {
		mockRequest.SetHeaders(headers)
	}

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(222)
	mockResponse.SetBody("i'm a cute loving mock, almost as cute as mumi, bichi and rasti")

	mock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	return mock
}
