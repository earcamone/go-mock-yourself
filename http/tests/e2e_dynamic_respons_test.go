package end_to_end

import (
	"fmt"
	"testing"
	"net/http"
	"io/ioutil"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http"
	"github.com/earcamone/go-mock-yourself/http/helpers"
)

//
// TestDynamicResponse() will ensure the Dynamic Responses constructor builds a proper Dynamic Mock Response
//

func TestDynamicResponse(t *testing.T) {
	const totalResponses = 30

	//
	// Build Dynamic Mock Responses with Dynamic Response content
	//

	var responses []*go_mock_yourself_http.Response

	for i := 0; i < totalResponses; i++ {
		//
		// Generate "random" content to build our "random" Mock Response
		//

		loopString := fmt.Sprintf("response%d", i)

		headers := make(map[string]string)
		headers[loopString] = loopString

		//
		// Build Mock Response with "random" content
		//

		mockResponse := new(go_mock_yourself_http.Response)

		if i % 5 == 0 {
			mockResponse.SetError(fmt.Errorf(loopString))

		} else {
			mockResponse.SetBody(loopString)
			mockResponse.SetHeaders(headers)
			mockResponse.SetStatusCode(http.StatusOK + i)
		}

		fmt.Println(mockResponse)

		//
		// Save responses in array for later assertions
		//

		responses = append(responses, mockResponse)
	}

	//
	// Build Dynamic Response
	//

	dynamicResponse := go_mock_yourself_http.NewDynamicResponse(responses...)

	//
	// Build always matching Mock with our Dynamic Response and install it :P
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, dynamicResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Loop through Mock Responses to ensure our Dynamic Response is returning the corresponding content
	//

	ass := assert.New(t)
	mockName := "bogus mock name"

	//
	// Let's iterate over the responses over and over again to ensure the Dynamic Response
	// is getting back to its first Mock Response after the last Mock Response is served :P
	//

	responses = append(responses, responses...)
	responses = append(responses, responses...)
	responses = append(responses, responses...)
	responses = append(responses, responses...)

	for responseIndex, mockResponse := range responses {
		response, responseError := http.Get("http://notengoenie.com")

		if responseIndex % 5 == 0 {
			ass.Nil(response)
			ass.Equal(mockResponse.GetError(mockName, new(http.Request)), responseError)

		} else {
			ass.Nil(responseError)

			responseHeaders := go_mock_yourself_helpers.HTTPHeadersToMap(response.Header)
			responseStatusCode := response.StatusCode
			responseBody, _ := ioutil.ReadAll(response.Body)

			ass.Equal(mockResponse.GetStatusCode(mockName, new(http.Request)), responseStatusCode)
			ass.Equal(mockResponse.GetHeaders(mockName, new(http.Request)), responseHeaders)
			ass.Equal(mockResponse.GetBody(mockName, new(http.Request)), string(responseBody))
		}
	}
}

