package go_mock_yourself_http

import (
	"os"
	"fmt"
	"net/http"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http/helpers"
)

func dumpCommunicationStream(reqonse interface{}) []byte {
	var turururu []byte

	switch reqonse.(type) {
	case *http.Request:
		request := reqonse.(*http.Request)

		//
		// Decorate missing headers with its default values as Go's http package would do before sending requests
		//

		body, _ := go_mock_yourself_helpers.GetBody(request)
		headers := go_mock_yourself_helpers.DecorateDefaultHeaders(request, go_mock_yourself_helpers.HTTPHeadersToMap(request.Header))

		//
		// Build HTTP request text representation
		//

		turururu = []byte(dumpRequestHeader(request) + "\r\n")
		turururu = append(turururu, []byte(dumpHeaders(headers) + "\r\n")...)
		turururu = append(turururu, body...)
		turururu = append(turururu, []byte("\r\n")...)

	case *http.Response:
		response := reqonse.(*http.Response)

		body, _ := go_mock_yourself_helpers.GetBody(response)
		headers := go_mock_yourself_helpers.HTTPHeadersToMap(response.Header)

		//
		// Build HTTP response text representation
		//

		turururu = []byte(dumpResponseHeader(response) + "\r\n")
		turururu = append(turururu, []byte(dumpHeaders(headers) + "\r\n")...)
		turururu = append(turururu, body...)
		turururu = append(turururu, []byte("\r\n")...)
	}

	return turururu
}

func dumpRequestHeader(request *http.Request) string {
	return fmt.Sprintf("%s %s %s", request.Method, request.URL.String(), request.Proto)
}

func dumpResponseHeader(response *http.Response) string {
	return fmt.Sprintf("%s %d %s", response.Proto, response.StatusCode, http.StatusText(response.StatusCode))
}

func dumpHeaders(headers map[string]string) string {
	stringHeaders := ""

	for header, value := range headers {
		stringHeaders += fmt.Sprintf("%s: %s\r\n", header, value)
	}

	return stringHeaders
}

func debug(format string, a ...interface{}) {
	if os.Getenv(envDebugging) == "on" {
		fmt.Printf(format, a...)
	}
}
