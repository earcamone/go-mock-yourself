package end_to_end_test

import (
	"time"
	"testing"
	"net/http"

	"github.com/stretchr/testify/assert"

	"github.com/mercadolibre/go-mock-yourself/http/tests/internal/e2e_helpers"
	"github.com/mercadolibre/go-mock-yourself/http"
)

//
// TestURLRegexMatchingMock() will ensure Mock Request URL Regex Matching scheme is working properly
//

func TestMockTimeout(t *testing.T) {
	ass := assert.New(t)

	//
	// Request timeout constant
	//

	const requestsTimeout = time.Second * 5

	//
	// Loop through all supported methods and ensure Timeout scheme is working like a charm
	//

	for _, method := range e2e_helpers.SupportedHTTPMethods {
		//
		// Reset Installed Mocks
		//

		go_mock_yourself_http.Reset()

		//
		// Generate URL Matching Mock
		//

		mock := e2e_helpers.CreateMock(nil, ".*", nil, nil)
		mock.Timeout = time.Second * 5

		//
		// Generate HTTP request matching just generated mock
		//

		request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", nil, "")

		//
		// Install matching Mock
		//

		go_mock_yourself_http.Install(mock)

		//
		// Issue request and calculate its processing timeout
		//

		now := time.Now()
		http.DefaultClient.Do(request)
		requestTimeout := time.Since(now)

		//
		// Ensure processing time equals at least the Request set Timeout
		//

		ass.True(requestTimeout >= requestsTimeout, "Mock timeout seems not to be working. Expected timeout: %d, processing timeout: %d, method: %s", requestsTimeout, requestTimeout, method)
	}
}
