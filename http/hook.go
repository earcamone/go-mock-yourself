package go_mock_yourself_http

import (
	"time"
	"net/http"
)

//
// goMockYourselfDo() is Go Mock Yourself's native HTTP Package Do() method replacement/hook method, in other
// words is where all the magic happens, similar to the hat of a magician but without the dead bird, for now..
//

func goMockYourselfDo(self *http.Client, request *http.Request) (response *http.Response, responseError error) {
	//
	// Get current request time, we might use this information at the end of the request processing
	// to sleep X time if client requested a specific processing Timeout for the current request
	//

	requestTime := time.Now()

	//
	// Current HTTP request matches an installed Mock?
	//

	mock := getMatchingMock(request)

	//
	// Should we log current stream communication?
	//

	matchingMock := mock != nil
	loggingSchemeOn := matchingMock && mock.Logging != nil

	if loggingSchemeOn {
		mock.Logging.Write(dumpCommunicationStream(request))
		mock.Logging.Flush()
	}

	//
	// Mock might match BUT should we mock it or let it proceed to target server? lets cut the logic in easy pieces
	//
	// NOTE: Mocks without available Response are usually used for logging purposes, you simply build a Mock with
	// some sort of matching criteria, add the optional logging stream and set no Response making Go Mock Yourself
	// only log the stream and redirect the request to Go's native HTTP package.
	//

	availableMockResponse := matchingMock && mock.Response != nil
	shouldMock := availableMockResponse && mock.Request.ShouldMock(request)

	//
	// Process HTTP request
	//

	switch shouldMock {
	case false:
		//
		// Either there is no Mock Response for the current Mock (technique usually used either to just log specific
		// URLs communication stream or to get a registered ShouldMock() callback called) or current request does not
		// match any registered Mock at all, either case lets call Go's client instance native Do method
		//

		hookUninstall()
		response, responseError = self.Do(request)
		hookInstall()

	case true:
		//
		// Add HTTP Client Request Cookies if available
		//
		// NOTE: code friendly stolen from http package
		//

		if self.Jar != nil {
			for _, cookie := range self.Jar.Cookies(request.URL) {
				request.AddCookie(cookie)
			}
		}

		//
		// Get Mocking Response
		//

		response, responseError = createHttpResponseFromMock(request, *mock)

		//
		// NOTE: Internally, Go native client's Do() method Close the request body and as we
		// are hooking it and not calling it in turn, we close it to prevent resource leaks.
		//
		// WARNING: This code was developed in an afternoon and the approach was to get a cool
		// and functional interface quick SO we might be missing some internal functionality
		// the Do() method might be doing too, this stuff will get resolved with time but the
		// warning is still here so you are aware just in case you start seeing some sort of leak.
		//
		// Still, this tool is intended for debugging purposes or tests development so you
		// should be safe :P
		//

		if request.Body != nil {
			request.Body.Close()
		}

		//
		// Set client new Cookies
		//
		// NOTE: code friendly stolen from http package
		//

		if self.Jar != nil {
			if rc := response.Cookies(); len(rc) > 0 {
				self.Jar.SetCookies(request.URL, rc)
			}
		}
	}

	//
	// Should we log current stream communication?
	//

	if loggingSchemeOn {
		mock.Logging.Write(dumpCommunicationStream(response))
		mock.Logging.Flush()
	}

	//
	// Does response require a specific processing timeout?
	//

	requestProcessingTime := time.Since(requestTime)

	if matchingMock && mock.Timeout != 0 && mock.Timeout - requestProcessingTime > 0 {
		time.Sleep(mock.Timeout - requestProcessingTime)
	}

	//
	// Return response processing, finally!!
	//

	return response, responseError
}
