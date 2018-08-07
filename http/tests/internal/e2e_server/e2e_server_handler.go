package e2e_server

import (
	"time"
	"math/rand"

	// Third-Party Imports
	"github.com/gin-gonic/gin"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
	"sync"
)

var e2eServerResponse *Response
var e2eServerResponseMutex sync.Mutex

func e2eTestingServerHandler(method string) func(c *gin.Context) {
	e2eHandler := func(c *gin.Context) {
		//
		// Generate Random Response Information
		//

		response := Response {
			Code: randomStatusCode(),
			Body: go_mock_yourself_helpers.RandomString(100),
			Headers: randomHeaders(),
		}

		//
		// Send Random Response
		//

		for header, value := range response.Headers {
			c.Header(header, value)
		}

		//
		// Save returned response for later test cases assertions using GetResponse()
		//

		SaveResponse(&response)

		//
		// Send response :P
		//

		c.Data(response.Code, "text/html", []byte(response.Body))
	}

	return e2eHandler
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func randomStatusCode() int {
	return random(200, 510)
}

func randomHeaders() map[string]string {
	randomString1 := go_mock_yourself_helpers.RandomString(100)
	randomString2 := go_mock_yourself_helpers.RandomString(100)

	headers := make(map[string]string)
	headers[randomString1] = randomString2

	return headers
}