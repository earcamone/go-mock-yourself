package go_mock_yourself_response

import (
	"net/http"
)

//
// DynamicResponse() will build a Dynamic HTTP Mock Response that will return the received Mock Responses
// in the order they were received. When the last response content is returned, if subsequent incoming
// requests match the Mock, the returned responses content loop will restart.
//

func DynamicResponse(responses ...*Response) *Response {
	//
	// Build "Dynamic Looping Response"
	//

	loopingCount := -1
	loopingResponse := new(Response)

	//
	// NOTE: Internally, Go Mock Yourself always calls first the GetError() method of a Matching Mock Response
	// while attempting to build the Mock HTTP Response thus incrementing the current dynamic response loop in
	// the Dynamic Response GetError() method would allow us to properly index the corresponding loop Response
	//
	// PD: The previous note sucks, this note will be improved in the official release, or not..
	//

	loopingResponse.SetError(func(mockName string, request http.Request) error {
		loopingCount++

		if loopingCount == len(responses) {
			loopingCount = 0
		}

		return responses[loopingCount].GetError(mockName, &request)
	})

	loopingResponse.SetStatusCode(func(mockName string, request http.Request) int {
		return responses[loopingCount].GetStatusCode(mockName, &request)
	})

	loopingResponse.SetHeaders(func(mockName string, request http.Request) map[string]string {
		return responses[loopingCount].GetHeaders(mockName, &request)
	})

	loopingResponse.SetBody(func(mockName string, request http.Request) string {
		return responses[loopingCount].GetBody(mockName, &request)
	})

	return loopingResponse
}
