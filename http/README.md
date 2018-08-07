
# Go Mock Yourself HTTP

[![Build Status](https://travis-ci.org/earcamone/go-mock-yourself.svg?branch=develop)](https://travis-ci.org/earcamone/go-mock-yourself)
[![codecov](https://codecov.io/gh/earcamone/go-mock-yourself/branch/develop/graph/badge.svg)](https://codecov.io/gh/earcamone/go-mock-yourself)

## Summary

Go Mock Yourself HTTP package will allow you (hopefully) to easily mock any complex HTTP communication flow due to its
flexible and robust schemes without having to change a single line of code of your application. Go Mock Yourself HTTP
package replaces at runtime Go's native http package functions in order to make it behave like you wish, taking the 
overhead of having to develop complex test code to mimik your production environment. 

Matching incoming HTTP requests to serve your Mock HTTP responses can be either accomplished using regular expressions 
on every aspect of an incoming HTTP request (method, url, headers, body and more) or you can even register what we call 
"Dynamic Matching Functions" that will allow you at run-time to determine if an incoming request should match or not
a specific mock (you could for example connect to a database and based on its information mock or not the request).

Mock HTTP Responses can also be built dynamically at run-time or you can even make any HTTP call (for example http.Post)
just fail with an also dynamically generated (or not) Go error.

Additionally, Go Mock Yourself HTTP package allows you to do lots of other cool things such as (among many others):

 * Pausing/Playing Go Mock Yourself HTTP Mocking Scheme at run-time
 
 * Forcing HTTP requests to last a specific duration (no matter if you decide to mock the response or allow it to 
   reach the target server)
   
 * Just plug it to your application and log specific communication flows in a nice human readable format to debug
   complex scenarios or even build mocks for a different mocking solution ;)
   
 * Conditionally (and dynamically) serve a Mock Response or allow the request to be sent to the remote server

## Table Of Content

 - [Introduction](#introduction)
 - [How it works?](#how-it-works)
 - [Matching Requests](#matching-requests)
 - [Matching Requests Priority](#matching-requests-priority)
 - [Mocking Responses](#mocking-responses)
 - [Installation](#installation)
 - [Plugging it in your application](#plugging-it-in-your-application)
 - [Matching Requests by HTTP Method](#matching-requests-by-http-method)
 - [Matching Requests by HTTP URL](#matching-requests-by-http-url)
 - [Matching Requests by HTTP Headers](#matching-requests-by-http-headers)
 - [Matching Requests by HTTP Body](#matching-requests-by-http-body)
 - [Matching Complex / Non Supported Requests Criterias](#matching-complex-non-supported-requests-criterias)
 - [Mocking Dynamic HTTP Responses](#mocking-dynamic-http-responses)
   - [Mocking Dynamic HTTP Response Codes](#mocking-dynamic-http-response-codes)
   - [Mocking Dynamic HTTP Response Headers](#mocking-dynamic-http-response-headers)
   - [Mocking Dynamic HTTP Response Body](#mocking-dynamic-http-response-body)
   - [Mocking Dynamic HTTP Response Failures](#mocking-dynamic-http-response-failures)
 - [Pausing Go Mock Yourself Scheme](#pausing-go-mock-yourself-scheme)
 - [Playing Go Mock Yourself Scheme](#playing-go-mock-yourself-scheme)
 - [Removing Installed Mocks](#removing-installed-mocks)
 - [Making HTTP Requests Last Specific Durations](#making-http-requests-last-specific-durations)
 - [Logging Specific HTTP Streams](#logging-specific-http-streams)
 - [Advanced Techniques](#advanced-techniques)
   - [To Mock Or Not To Mock?](#to-mock-or-not-to-mock)
   - [Asserting Incoming Requests Information Server-Side](#asserting-incoming-requests-information-server-side)
   - [Registering N different Responses for subsequent N Requests in the same matching Mock](#registering-n-different-responses-for-subsequent-n-requests-in-the-same-matching-mock)
   - [Serving a Mock Response only in the N request](#serving-a-mock-response-only-in-the-n-request)

## Introduction

Mocking HTTP responses is so easy (hopefully) that we will approach each package feature by example. Still, there are
some quick basic notes you should have in mind while inspecting the package which we will cover in the following sections
prior to each scheme example BUT before that lets all do a huge pogo (jijiji) for "Bouke van der Bijl" and its amazing 
["Monkey Patching"](https://github.com/bouk/monkey) package which makes this project possible! If you like this package
kindly go to its repo and show some love to it too!

`NOTE:` If you feel brave enough, just skip the following sections and go directly to each scheme example. 

## How it works?

Go Mock Yourself HTTP intercepts at run-time every call to Go's native http package and mimiks its behaviour in order to
serve the Mock Responses you might install. This means that after importing Go Mock Yourself HTTP to any application,
from there on all HTTP requests done by the application (provided they are done through Go's native http package) will 
be processed by Go Mock Yourself and if the request matches any of your installed Mocks, it will automatically serve
them without allowing the request reaching Go's package. On the other hand, if the request does not match any installed
Mock it will simply proxy the request to Go's http package allowing the application to reach the remote server.

`WARNING:` There are some warnings you must have in mind when using the package so kindly read them at the bottom of the
documentation in the [DISCLAIMER section](#disclaimer).

## Matching Requests

Go Mock Yourself supports the use of Regular Expressions in all its supported Matching Criterias:
 
 * Request Method
 * Request URL
 * Request Headers
 * Request Body
 
This means you can register a Mock to match any of the previously mentioned criterias individually or grouped, for 
example only matching requests to a given URL, containing a specific header and a specific word in its body. 
 
Additionally it supports Dynamic Matching Functions for each criteria which you can register to for example mock only 
the third request to a same URL. Dynamic Matching Functions are a powerful feature that will allow you to build complex
matching criterias when using regular expressions is not enough for you, for example using counters or even comparing
requests information against a database in order to asses if Go Mock Yourself should return or not a Mock Response.

## Matching Requests Priority

Whenever an incoming HTTP request is processed, each installed Mock is applied in the order they were installed to asses 
if the request should be served an installed Mock Response or not. For example, you could first install a series of Mocks 
to serve specific Mock Responses for specific HTTP requests and at the end of these Mocks simply install a Mock that 
would match ANY url to serve a generic not found Mock Response for those incoming requests that didn't match any of the 
previously installed Mocks.

## Mocking Responses

Mock responses can be built either statically or dynamically, allowing you to for example retrieve information from
a database to build your Mock Responses at run-time.

Mocks MUST ALWAYS have a Matching Mock Request BUT they can be missing a Mock Response as you might just want to for 
example use a Matching Mock to intercept specific requests to assert the corresponding information is being sent to a 
remote server BUT still allow the remote server to process the request instead of your installed mock to assert the
response from this service is correct.

In resume, whenever a Mock lacks a Mock Response all its matching criterias will be applied and if its a match, the
request will be sent to the target URL without Go Mock Yourself interfering with it (no further installed Mocks will
be attempted to match the request).

## Installation

```
go get github.com/earcamone/go-mock-yourself/http
```

## Plugging it in your application

Go Mock Yourself will start intercepting your application requests whenever its imported so we recommend the following
recipe to easily plug it to your application without having to change a single line of code, just create a file called 
go_mock_yourself.go in your application scaffolding and dump in it the following code:

```
import (
    "github.com/earcamone/go-mock-yourself/http"
)

func init() {
    go_mock_yourself_http.Play()
}
```

## Matching Requests by HTTP Method

### Using Regular Expressions

The following Mock will match any GET or POST request:

```
    const mySweetMockBody = "i'm a happy and loving mock response"
    
    //
    // Build Matching Mock
    //
    
    mockRequest := new(go_mock_yourself_http.Request)
    mockRequest.SetMethod(`(GET|POST)$`)

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
	
	response, _ := http.Get("http://notengoenie.com")
	body, _ := ioutil.ReadAll(response.Body)

    if string(body) == mySweetMockBody {
	    fmt.Println("lovely mock, you are all i need, would you merry me?")
    }
```

### Using Dynamic Matching Function

The following Mock will match only the third GET request since your application is started using a Dynamic Matching 
Function with a simple counter:

```
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
```

`NOTE:` You could make the request match your current mock without even analyzing the request method, when it comes to
Dynamic Matching Functions its up to you (actually your functions return value) making the mock match or not the current
HTTP request with whatever criteria you choose, you just get the request method as a parameter but if you decide to
ignore it and just return always true is up to you. 

## Matching Requests by HTTP URL

### Using Regular Expressions

The following Mock will match requests sent to any URL:

```
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
```

### Using Dynamic Matching Function

The following Mock will match requests sent to any URL:

```
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
```

## Matching Requests by HTTP Headers

### Using Regular Expressions

The following Mock will match any request containing an "X-Request-Id" header (cof cof [ginrequestid](https://github.com/atarantini/ginrequestid) cof cof):

```
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
```

NOTE: You can also use regular expressions to match the Headers names, not only their values ;)

### Using Dynamic Matching Function

The following Mock will match any request containing an "X-Request-Id" header (cof cof [ginrequestid](https://github.com/atarantini/ginrequestid) cof cof):

```
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
```

## Matching Requests by HTTP Body

### Using Regular Expressions

The following Mock will match any request which body contains the phrase "linters suck":

```
	const mySweetMockBody = "eventhough bubi likes using linters and middlewares in excess, i still love him"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetBody(`.*linters suck.*`)

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

	requestBody := strings.NewReader("i understand the arguments behind it, still linters suck.")
	response, _ := http.Post("http://supercalifragilisticoespialidoso.com", "text/html", requestBody)

	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
```

### Using Dynamic Matching Function

The following Mock will match any request with a body containing the phrase "linters suck":

```
	const mySweetMockBody = "eventhough bubi likes using linters, i still love him"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
    mockRequest.SetBody(func(requestBody string) bool {
        return strings.Contains(requestBody, "linters suck")
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

	requestBody := strings.NewReader("i understand the arguments behind it, still linters suck.")
	response, _ := http.Post("http://supercalifragilisticoespialidoso.com", "text/html", requestBody)

	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
```

## Matching Complex / Non Supported Requests Criterias

Sometimes you might have complex scenarios were specific matching criterias (for example the body content) might be
a matching mock or not depending on multiple other criterias values and building this type of matching mocks with all 
its logic splited in different Dynamic Matching Functions might become a mess. You might even want to build a matching 
Mock only if the HTTP version or other request information not listed in the current supported criterias has a specific 
value, don't worry, for these scenarios you can use the SetRequestCriteria() function which allows you to register a
Request Dynamic Matching Function.

The Request Dynamic Matching Function will receive directly the http.Request object representing the incoming request 
and depending on its return value, the Mock will become a match or not. 

Be adviced that when registering a Request Dynamic Matching Function, the Mock will simply match or not depending on
this function return value, NO FURTHER MATCHING CRITERIAS WILL BE APPLIED meaning that if the function returns false the
next installed mock will be applied and if it returns true it will simply handle the incoming request.

```
	const mySweetMockBody = "En carnavales de señales no verbales fue descubriendo el lenguaje del inconsciente en busca de alguien que lo pueda ver a través del follaje. Interpretó modestos gestos que en sí mismo vio y comprendió el mensaje. vacuna para incongruentes, se paró y gritó: bendito aprendizaje. Y de pronto sintió que se le inflaba el pecho, vertiginosa sensación. Entre ilusiones y comparaciones enjuició toda una vida entera. Y hoy ve como un juicio que antes servía, hoy no sirvió. Ayer si, hoy cualquiera. Pero ahora ¿cómo se hace, cómo saco esto de acá? ¿Cómo empiezo de nuevo? ¿Cómo perdono? ¿Cómo me perdono a mí además? ¿Cómo disfruto el juego? Y de pronto sintió un nudo en la garganta y sin embargo disfrutó. Él le llamó aceptación a ese llanto sin consuelo y desde ahí transformó la rigidez del miedo cruel y paralizador en impulso motor. Fue en busca de su esencia una y mil veces y encontró que ésta siempre mutaba, de forma espacios, tiempos, todo acorde a la emoción del momento en que estaba focalizó tanto en ahora que temió perder completa la memoria. Fue entonces que se hizo conciencia y creyó comprender: mi esencia no es mi historia, no Y de pronto sintió muy livianos los hombros y rumbo al cielo se cayó. Él le llamó plenitud a esa risa en carcajada y desde ahí la virtud de vivir libre o nada creció. Como un alud eligió ver la luz. Él le llamó aceptación a ese llanto sin consuelo y desde ahí transformo la rigidez del miedo cruel y paralizador en impulso motor. Él le llamó plenitud a esa risa en carcajada y desde ahí la virtud de vivir libre o nada creció"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetRequestCriteria(func(request *http.Request) bool {
		_, url, _, body := go_mock_yourself_helpers.RequestInformation(request)
		return strings.Contains(url, "?param1") && strings.Contains(body, mySweetMockBody)
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

	requestBody := strings.NewReader(mySweetMockBody)
	response, _ := http.Post("http://supercalifragilisticoespialidoso.com/?param1=bleh", "text/html", requestBody)

	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
```

## Mocking Dynamic HTTP Responses

In the previous sections we went through almost all Go Mock Yourself HTTP Requests matching schemes but we always returned 
the same static Mock Response, in this section we will cover in detail all Go Mock Yourself available Responses schemes.

### Mocking Dynamic HTTP Response Codes

The following Mock will return a dynamically generated HTTP Status Code:

```
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
```

### Mocking Dynamic HTTP Response Headers 

The following Mock will return dynamically generated HTTP Headers:

```
	const mySweetMockBody = "even if he tried it very hard, there is a word coli will never be able to stop saying"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	mockResponse.SetHeaders(func(matchingMockName string, request http.Request) map[string]string {
		headers := make(map[string]string)
		headers["Coli-Is"] = "everybody knows.."

		return headers
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

	if string(body) == mySweetMockBody {
		for header, value := range response.Header {
			if header == "Coli-Is" && value[0] == "everybody knows.." {
				fmt.Println("lovely mock, you are all i need, would you merry me?")
			}
		}
	}
```

### Mocking Dynamic HTTP Response Body 

The following Mock will return a dynamically generated HTTP Body:

```
	const mySweetMockBody = "No te quedes inmóvil al borde del camino no congeles el júbilo no quieras con desgana no te salves ahora ni nunca no te salves no te llenes de calma no reserves del mundo sólo un rincón tranquilo no dejes caer los párpados pesados como juicios no te quedes sin labios no te duermas sin sueño no te pienses sin sangre no te juzgues sin tiempo, pero si pese a todo no puedes evitarlo y congelas el júbilo y quieres con desgana y te salvas ahora y te llenas de calma y reservas del mundo sólo un rincón tranquilo y dejas caer los párpados pesados como juicios y te secas sin labios y te duermes sin sueño y te piensas sin sangre y te juzgas sin tiempo y te quedas inmóvil al borde del camino y te salvas, entonces no te quedes conmigo"

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	
	mockResponse.SetBody(func(matchingMockName string, request http.Request) string {
        return mySweetMockBody
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

	if string(body) == mySweetMockBody {
        fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
```
	
### Mocking Dynamic HTTP Response Failures 

The following Mock will make any HTTP request fail with a dynamic generated error:

```
	const mySweetMockError = "Si para recobrar lo recobrado debí perder primero lo perdido, si para conseguir lo conseguido tuve que soportar lo soportado, si para estar ahora enamorado fue menester haber estado herido, tengo por bien sufrido lo sufrido, tengo por bien llorado lo llorado. Porque después de todo he comprobado que no se goza bien de lo gozado sino después de haberlo padecido. Porque después de todo he comprendido por lo que el árbol tiene de florido vive de lo que tiene sepultado."

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetError(func(matchingMockName string, request http.Request) error {
		return fmt.Errorf(mySweetMockError)
	})

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	_, responseError := http.Get("http://notengoenie.com")

	if responseError.Error() == mySweetMockError {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
```

## Pausing Go Mock Yourself Scheme

There might be scenarios where suddenly you need the Mocking Scheme to temporarily stop working until a specific
condition is triggered, when this happens you can easily Pause Go Mock Yourself HTTP Mocking scheme using Pause()

```
    go_mock_yourself_http.Pause()
```

## Playing Go Mock Yourself Scheme

You can activate again at any given time Go Mock Yourself Mocking scheme using Play() function:

```
    go_mock_yourself_http.Play()
```

NOTE: For more information regarding pausing Go Mock Yourself HTTP Mocking scheme kindly check [Pausing Go Mock Yourself Scheme](#playing-go-mock-yourself-scheme)
 
## Removing Installed Mocks

There might be scenarios were suddenly you might need to reconfigure all your installed Mocks at run-time, in this
scenario you can easily call Go Mock Yourself HTTP Reset() function which will remove all currently installed mocks:

```
    go_mock_yourself_http.Reset()
```

## Making HTTP Requests Last Specific Durations

Whatever HTTP response you make your matching Mock return (even errors or those redirected to a remote server), you can
make them last the duration you want:

```
	const mySweetMockBody = "Vendrá la muerte y tendrá tus ojos —esta muerte que nos acompaña de la mañana a la noche, insomne, sorda, como un viejo remordimiento o un vicio absurdo. Tus ojos serán una palabra hueca, un grito ahogado, un silencio. Así los ves cada mañana cuando a solas te inclinas hacia el espejo. Oh querida esperanza, ese día también sabremos que eres la vida y la nada. Para todos tiene la muerte una mirada. Vendrá la muerte y tendrá tus ojos. Será como dejar un vicio, como mirar en el espejo asomarse un rostro muerto, como escuchar un labio cerrado. Nos hundiremos en el remolino, mudos."

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)
	mySweetMock.Timeout = time.Second * 5

	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	now := time.Now()
	response, _ := http.Get("http://notengoenie.com")
	elapsed := time.Since(now)

	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Printf("lovely mock, you are all i need eventhough it took you %v seconds to finish, would you merry me?\n", elapsed)
	}
```

Note how the Timeout option is attached to the Mock itself and not the Mock Response as you can install a matching mock
without Mock Response which would make the request be redirected directly to the target remote server but still, once
the answer is back, make it last the specified duration.

## Logging Specific HTTP Streams

Let's say you are having some troubles debugging your current production code and you want to easily know what is exactly 
being sent to a remote server and what is being received, you can easily install a Matching Mock that will only log the
HTTP stream and do nothing else:

```
	const mySweetMockBody = "Lo de menos son todos los secretos que intuyo, huelo, toco y siempre te respeto. Lo de menos es que jamás me sobres, que tu amor me enriquezca haciéndome más pobre. Lo de menos es que tus sentimientos no marchen en horario con mi renacimiento. Lo de menos es larga soledad, lo de menos es cuánto corazón. Lo que menos importa es mi razón lo de menos incluso es tu jamás, mientras cante mi amor intentando atrapar las palabras que digan lo de más. Amoroso, de forma que no mancha, en verso y melodía recurro a la revancha. Mi despecho de besará la vida allá donde más sola o donde más querida. Dondequiera que saltes o que gires habrá un segundo mío para que lo suspires. Es la prenda de larga soledad, es la prenda de cuánto corazón. Pajarillo, delfín de mis dos rosas, espántame los golpes y no la mariposa. Ejercita tu danza en mi cintura, aroma incomparable, oh pan de mi locura. Con tu cuerpo vestido de mis manos haré una nueva infancia, al borde del océano. Desde el mar te lo cuento en soledad, desde el mar te lo lanza un corazón."

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mockResponse := new(go_mock_yourself_http.Response)
	mockResponse.SetStatusCode(http.StatusOK)
	mockResponse.SetBody(mySweetMockBody)

	//
	// Install Mock
	//

	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)

	file, _ := os.Create("/tmp/gomockyourself.txt")
	mySweetMock.Logging = bufio.NewWriter(file)

	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	http.Get("http://notengoenie.com")
	loggingData, _ := ioutil.ReadFile("/tmp/gomockyourself.txt")

	if strings.Contains(string(loggingData), mySweetMockBody) {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
```

NOTE: In the example we actually installed a Mock Response as we want the examples to work even if you have no internet
connection BUT you could simply remove the three mockResponse lines and instead of logging the Mock Response, the actual
remote server response would be logged.
  
## Advanced Techniques

### To Mock Or Not To Mock?

Let's say there are conditions at which a specific installed Mock should not match incoming requests, there is a 
callback you can register called ShouldMock that allows a Matching Mock to be skipped or not:

```
	const mySweetMockBody = "Que días más intensos estos no se pa' ti, pero pa' mi.. cada instante escarba más adentro. que días más intensos estos. Mis ojos no cierran, esperan, no guardan descanso. Esperan que traigan sorpresa los días, pues traigo un par de sueños acá anhelando ser vida. No quiero ser un tornillo más en la máquina de demoler. No quiero ser un soldado más en la guerra del poder. quiero reencontrar la inocencia y la pureza tal vez. Y si es posible querer sin ver a quién y sin saber por qué. Me vendría muy bien, nos vendría muy bien. Que días más intensos estos. no se pa' ti, pero pa' mí cada instante escarba más adentro. Que días más intensos estos. mi alma se escapa, se entrega. No guarda descanso. Espera que traigan mas reto los días, pues las cicatrices son mapas en mi piel Que orientan mi vida. No quiero estar un minuto más sin cultivar algún saber. No quiero ver otro día llegar sin crecer porque hoy no es lo mismo que ayer. Ya silvio me sumó su estatura con su canción de imágenes. Bob Marley me enseñó que no hay fortuna Que pague lo que ha logrado él. Hervid Hancock me enseñó a ser más fiel A lo que pide la piel y a no temerle a crecer Que días más intensos estos tanto afuera como adentro... "

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	mockRequest.SetShouldMock(func(request *http.Request) bool {
		_, _, _, body := go_mock_yourself_helpers.RequestInformation(request)
		return strings.Contains(body, "Que días más intensos estos")
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

	response, _ := http.Post("http://notengoenie.com", "text/html", strings.NewReader(mySweetMockBody))
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) == mySweetMockBody {
		fmt.Println("lovely mock, you are all i need, would you merry me?")
	}
```

In the previous example we installed an always matching mock to any targeted URL BUT we installed a ShouldMock callback
which ensures the body contains a specific string in order to match :)

What's the use of this scheme? well, imagination is your limit BUT for example there is an usual pattern where you just
want to Mock the third request to a specific url.. you could then just install a matching Mock for that URL and install
a ShouldMock function with a counter and only when it gets to the third request make it a match :)

### Asserting Incoming Requests Information Server-Side

If you develop good end-to-end test cases you should always ensure not only your expected response is there BUT that 
the response was generated after the corresponding request arrived server-side. You can easily do this with the 
ShouldMock() function explained in the previous section. 

Inside your ShouldMock() function you would assert that the matching incoming request information is correct and you 
would simply always return true so your Mock response is returned (or not, depending what you are trying to accomplish).

### Registering N different Responses for subsequent N Requests in the same matching Mock

Let's say there are scenarios were an exact same Matching Mock request requires to return different Mock Responses
based on the incoming request count, you can easily build this type of "Dynamic Mock Responses" using NewDynamicResponse
helper function:

```
	const mySweetMockBody = ""

	//
	// Build Matching Mock
	//

	mockRequest := new(go_mock_yourself_http.Request)
	mockRequest.SetUrl(".*")

	//
	// Build Dynamic Mock Response
	//

	mockResponse1 := new(go_mock_yourself_http.Response)
	mockResponse1.SetStatusCode(http.StatusOK)
	mockResponse1.SetBody("response1")

	mockResponse2 := new(go_mock_yourself_http.Response)
	mockResponse2.SetStatusCode(http.StatusCreated)
	mockResponse2.SetBody("response2")

	mockResponse3 := new(go_mock_yourself_http.Response)
	mockResponse3.SetError(fmt.Errorf("response3"))

	responses := []*go_mock_yourself_http.Response {
		mockResponse1,
		mockResponse2,
		mockResponse3,
	}

	//
	// Install Mock
	//

	mockResponse := go_mock_yourself_http.NewDynamicResponse(responses...)
	mySweetMock, _ := go_mock_yourself_http.NewMock("my lovely testing mock", mockRequest, mockResponse)

	go_mock_yourself_http.Install(mySweetMock)

	//
	// Issue HTTP Request and expect our Sweet Mock Response
	//

	bogusRequest := new(http.Request)

	for _, mockResponse := range responses {
		response, err := http.Get("http://notengoenie.com")

		body := []byte("")
		if err == nil {
			body, _ = ioutil.ReadAll(response.Body)
		}

		if responseError := mockResponse.GetError("", bogusRequest); responseError != nil && responseError == err {
			fmt.Println("lovely mock, you are all i need, would you merry me?")

		} else if string(body) == mockResponse.GetBody("", bogusRequest) && response.StatusCode == mockResponse.GetStatusCode("", bogusRequest) {
			fmt.Println("lovely mock, you are all i need, would you merry me?")
		}
	}
```

### Serving a Mock Response only in the N request

Kindly check the foot notes at the [To Mock Or Not To Mock?]() section ;)

### Randomly Failing Requests

A simple helper function is being implemented as we speak to support this scheme..

### Randomly Mocking Requests

A simple helper function is being implemented as we speak to support this scheme..

### Making Requests Take Random Timeouts

A simple helper function is being implemented as we speak to support this scheme..

### HTTP Streams Recording / Playing

Being implemented as we speak..

## DISCLAIMER

Go Mock Yourself HTTP is built over "Bouke van der Bijl" amazing ["Monkey Patching"](https://github.com/bouk/monkey) package
which as it states in its documentation you should have the following warnings under consideration when using his 
(and thus, ours) package:

1. Monkey sometimes fails to patch a function if inlining is enabled. Try running your tests with inlining disabled, 
   for example: go test -gcflags=-l. The same command line argument can also be used for builds.
   
2. Monkey won't work on some security-oriented operating system that don't allow memory pages to be both write and 
   execute at the same time. With the current approach there's not really a reliable fix for this.
   
3. Monkey is not threadsafe. Or any kind of safe.

4. I've tested monkey on OSX 10.10.2 and Ubuntu 14.04. It should work on any unix-based x86 or x86-64 system.

Now, considerations strictly about Go Mock Yourself HTTP:

IF you ALWAYS mock ALL the HTTP requests done by your application, meaning NO request will ever be skipped by Go Mock
Yourself HTTP and sent to the target server, then Go Mock Yourself HTTP `IS THREAD SAFE`. 

Now, if for any reason one HTTP request is not handled by a Go Mock Yourself Mock, while this request is being handled
by Go's native package, Go Mock Yourself would had stepped out of the way so if there are other requests in other threads
being performed that should match an installed Mock, Go Mock Yourself HTTP would not be able to handle it allowing the
request to reach the target server.

Additionally to the previously mentioned scenario, Panics could be triggered due to a specific scenario were a skipped 
request (allowed to be processed by Go's native package) is finishing its processing (thus about to get Go's native 
package hooked again) while other threads are entering/coming back from Go's native package http requests processing. 

In resume, if you want to be safe, either:

1. Never use threads (or ensure your threads run in series and not concurrently)

2. Use threads BUT ensure all your HTTP requests have a Matching Mock serving their Mock Responses

3. If you have some requests that will be handled by a remote server, ensure no other threads attempt to make any sort
   of HTTP request until the remote server finishes processing the request and Go Mock Yourself HTTP package takes
   again control of any HTTP request.

If enough people show some love to the project, i will implement my own thread-safe Monkey Patching package, making
Go Mock Yourself HTTP thread-safe no matter what scenario you run it :)
 
## CHANGELOG

### `0.0.1` - DD-MM-YYYY - Emiliano Arcamone (earcamone@hotmail.com)

 - Go Mock Yourself HTTP: Initial stable version

## Developers

[Emiliano Arcamone](https://www.linkedin.com/in/emiliano-arcamone-727a52100/)

## Grateful with:

[Eduardo Acosta Miguens](https://github.com/eduacostam): For helping me deal with Go's military enforcing standards and
helping me avoid having to read boring standard stuff.
