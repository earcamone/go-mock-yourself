package go_mock_yourself_response

import (
	"fmt"
	"testing"
	"net/http"

	// Third-Party Imports
	"github.com/stretchr/testify/assert"
)

//
// TestDynamicResponse() will ensure the Dynamic Responses constructor builds a proper Dynamic Mock Response
//

func TestDynamicResponse(t *testing.T) {
	const totalResponses = 30

	//
	// Build Dynamic Mock Responses with Dynamic Response content
	//

	var responses []*Response

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

		mockResponse := new(Response)

		if i % 5 == 0 {
			mockResponse.SetError(fmt.Errorf(loopString))

		} else {
			mockResponse.SetBody(loopString)
			mockResponse.SetHeaders(headers)
			mockResponse.SetStatusCode(http.StatusOK + i)
		}

		//
		// Save responses in array for later assertions
		//

		responses = append(responses, mockResponse)
	}

	//
	// Build Dynamic Response
	//

	dynamicResponse := DynamicResponse(responses...)

	//
	// Loop through Mock Responses to ensure our Dynamic Response is returning the corresponding content
	//

	ass := assert.New(t)

	mockName := "bogus mock name"
	mockRequest := new(http.Request)

	//
	// Let's iterate over the responses over and over again to ensure the Dynamic Response
	// is getting back to its first Mock Response after the last Mock Response is served :P
	//

	responses = append(responses, responses...)
	responses = append(responses, responses...)
	responses = append(responses, responses...)
	responses = append(responses, responses...)

	for responseIndex, response := range responses {
		dynamicError := dynamicResponse.GetError(mockName, mockRequest)

		dynamicStatusCode := dynamicResponse.GetStatusCode(mockName, mockRequest)
		dynamicHeaders := dynamicResponse.GetHeaders(mockName, mockRequest)
		dynamicBody := dynamicResponse.GetBody(mockName, mockRequest)

		if responseIndex % 5 == 0 {
			ass.Equal(response.GetError(mockName, mockRequest), dynamicError)

		} else {
			ass.Equal(response.GetBody(mockName, mockRequest), dynamicBody)
			ass.Equal(response.GetHeaders(mockName, mockRequest), dynamicHeaders)
			ass.Equal(response.GetStatusCode(mockName, mockRequest), dynamicStatusCode)
		}
	}
}
