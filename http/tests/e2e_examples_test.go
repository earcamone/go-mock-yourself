package end_to_end_test

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http"
)

//
// ExampleMatchMethodByRegularExpression() matches requests by HTTP method using regular expressions
//

func ExampleMatchMethodByRegularExpression() {
	const mySweetMockBody = "i'm a happy and loving mock response"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetMethod(`(GET|POST)$`)

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	response, _ := http.Get("http://notengoenie.com")
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}

//
// ExampleMatchMethodDynamically() matches requests by HTTP method dynamically at run-time using a Dynamic Matching Function
//

func ExampleMatchMethodDynamically() {
	const mySweetMockBody = "i'm a happy and loving mock response"

	//
	// Build Matching Mock
	//

	requestsCount := 0

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetMethod(func(requestMethod string) bool {
		if requestMethod == http.MethodGet {
			requestsCount++
		}

		return requestsCount == 3
	})

	//
	// Build Matching Mock's Response
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	for i := 1; i < 5; i++ {
		response, _ := http.Get("http://notengoenie.com")
		body, _ := ioutil.ReadAll(response.Body)

		if i == 3 && string(body) == mySweetMockBody {
			fmt.Println("lovely mock, you are all i need, would you merry me?")
		}
	}
}

//
// ExampleMatchURLByRegularExpression() matches requests by HTTP URL using regular expressions
//

func ExampleMatchURLByRegularExpression() {
	const mySweetMockBody = "rasti loves mindcraft, morse code and olorin"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(`.*`)

	//
	// Build Matching Mock's Response
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	response, _ := http.Head("http://supercalifragilisticoespialidoso.com")
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}

//
// ExampleMatchURLDynamically() matches requests by HTTP URL dynamically at run-time using a Dynamic Matching Function
//

func ExampleMatchURLDynamically() {
	const mySweetMockBody = "rasti loves mindcraft, morse code and olorin"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(func(requestUrl string) bool {
		return true // will match all requests
	})

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	response, _ := http.Head("http://supercalifragilisticoespialidoso.com")
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}

//
// ExampleMatchHeadersByRegularExpression() matches requests by HTTP URL using regular expressions
//

func ExampleMatchHeadersByRegularExpression() {
	const mySweetMockBody = "mumi is the sweetest guy in the condado, even though i will ever be jealous of the Arpia!"

	//
	// Build Matching Mock
	//

	matchingHeaders := make(map[string]string)
	matchingHeaders["^X-Request-Id$"] = ".*"

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetHeaders(matchingHeaders)

	//
	// Build Matching Mock Response
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	requestHeaders := make(http.Header)
	requestHeaders["X-Request-Id"] = []string{ mySweetMockBody }

	request, _ := http.NewRequest(http.MethodPost, "http://notengoenie.com", nil)
	request.Header = requestHeaders

	response, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}

//
// ExampleMatchHeadersDynamically() matches requests by HTTP Headers dynamically at run-time using a Dynamic Matching Function
//

func ExampleMatchHeadersDynamically() {
	const mySweetMockBody = "mumi is the sweetest guy in the condado, even though i will ever be jealous of the Arpia!"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetHeaders(func(requestHeaders map[string]string) bool {
		for header := range requestHeaders {
			if header == "X-Request-Id" {
				return true
			}
		}

		return false
	})

	//
	// Build Matching Mock Response
	//

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	requestHeaders := make(http.Header)
	requestHeaders["X-Request-Id"] = []string{ mySweetMockBody }

	request, _ := http.NewRequest(http.MethodPost, "http://notengoenie.com", nil)
	request.Header = requestHeaders

	response, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}

//
// ExampleMatchBodyByRegularExpression() matches requests by HTTP Body using regular expressions
//

func ExampleMatchBodyByRegularExpression() {
	const mySweetMockBody = "eventhough bubi likes using linters, i still love him"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetBody(`.*linters suck.*`)

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	requestBody := strings.NewReader("i understand the arguments behind it, still linters suck.")
	response, _ := http.Post("http://supercalifragilisticoespialidoso.com", "text/html", requestBody)

	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}

//
// ExampleMatchBodyDynamically() matches requests by HTTP Body dynamically at run-time using a Dynamic Matching Function
//

func ExampleMatchBodyDynamically() {
	const mySweetMockBody = "eventhough bubi likes using linters, i still love him"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetBody(func(requestBody string) bool {
		return strings.Contains(requestBody, "linters suck")
	})

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	requestBody := strings.NewReader("i understand the arguments behind it, still linters suck.")
	response, _ := http.Post("http://supercalifragilisticoespialidoso.com", "text/html", requestBody)

	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}

//
// ExampleDynamicHTTPCodeResponse() will return a dynamic mock response with a dynamically generated HTTP code
//

func ExampleDynamicHTTPCodeResponse() {
	const mySweetMockBody = "nobody, and i mean N O B O D Y, can drink alcohol like elo"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetBody(mySweetMockBody)

	mockResponse.SetStatusCode(func(matchingMockName string, request http.Request) int {
		return 666
	})

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	response, _ := http.Get("http://notengoenie.com")
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody && response.StatusCode == 666 {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
}