package go_mock_yourself_helpers

import (
	"bytes"
	"strings"
	"net/http"
	"io/ioutil"
)

//
// RequestInformation() will return the received HTTP request information
//

func RequestInformation(request *http.Request) (string, string, map[string]string, string) {
	//
	// Read request body
	//

	body, _ := GetBody(request)

	//
	// Read headers and convert them to a simple map[string]string
	//

	headers := HTTPHeadersToMap(request.Header)

	//
	// Read target URL and Method
	//

	url := ""

	if request.URL != nil {
		url = request.URL.String()
	}

	method := request.Method

	return method, url, headers, string(body)
}

//
// ResponseInformation() will return the received HTTP request information
//

func ResponseInformation(response *http.Response) (int, map[string]string, string) {
	//
	// Read response body
	//

	body, _ := GetBody(response)

	//
	// Read headers and convert them to a simple map[string]string
	//

	headers := HTTPHeadersToMap(response.Header)

	//
	// Return HTTP Response Information
	//

	return response.StatusCode, headers, string(body)
}

//
// GetBody() will read the received HTTP Request/Response body
//

func GetBody(reqonse interface{}) ([]byte, error) {
	var body []byte
	var bodyError error

	switch reqonse.(type) {
	case *http.Request:
		request := reqonse.(*http.Request)

		if request != nil && request.Body != nil {
			body, bodyError = ioutil.ReadAll(request.Body)

			if bodyError == nil {
				request.Body.Close()

				bodyBuffer := bytes.NewBuffer(body)
				request.Body = ioutil.NopCloser(bodyBuffer)
			}
		}

	case *http.Response:
		response := reqonse.(*http.Response)

		if response != nil && response.Body != nil {
			body, bodyError = ioutil.ReadAll(response.Body)

			if bodyError == nil {
				response.Body.Close()

				bodyBuffer := bytes.NewBuffer(body)
				response.Body = ioutil.NopCloser(bodyBuffer)
			}
		}
	}

	return body, bodyError
}

//
// HTTPHeadersToMap() will build a Headers map[string]string from the received http.Headers
//

func HTTPHeadersToMap(headers http.Header) map[string]string {
	if headers == nil {
		return nil
	}

	responseHeaders := make(map[string]string)

	for header, values := range headers {
		responseHeaders[header] = strings.Join(values, "; ")
	}

	return responseHeaders
}

//
// HeadersMapToHTTPHeaders() will build an http.Header from a Headers map[string]string
//

func HeadersMapToHTTPHeaders(headers map[string]string) http.Header {
	responseHeaders := make(http.Header)

	for header, value := range headers {
		responseHeaders[header] = strings.Split(value, ";")
	}

	return responseHeaders
}

//
// DecorateDefaultHeaders() will decorate the received HTTP Request with those default
// HTTP headers Go's http package adds to any HTTP request if they are missing
//

func DecorateDefaultHeaders(request *http.Request, headers map[string]string) map[string]string {
	headersCopy := DuplicateMap(headers)

	//
	// Handle default HTTP Headers internally set by HTTP package before requests are sent
	//

	headersCopy["Client"] = headerClient

	if request.URL != nil {
		headersCopy["Host"] = request.URL.Hostname()
	}

	if _, connectionHeaderAvailable := headersCopy["Connection"]; connectionHeaderAvailable == false {
		headersCopy["Connection"] = "close"
	}

	noRangePresent := request.Header.Get("Range") == ""
	noEncodingPresent := request.Header.Get("Accept-Encoding") == ""

	if noEncodingPresent && noRangePresent && request.Method != "HEAD" {
		headersCopy["Accept-Encoding"] = "gzip"
	}

	return headersCopy
}
