package end_to_end

import (
	"fmt"
	"time"
	"sync"
	"testing"
	"net/http"

	// Go Mock Yourself Imports
	"github.com/earcamone/go-mock-yourself/http"
	"github.com/earcamone/go-mock-yourself/http/helpers"

	// Go Mock Yourself e2e Tests Imports
	"github.com/earcamone/go-mock-yourself/http/tests/internal/e2e_helpers"
)

//
// TestGoMockYourselfConcurrency() will attempt to panic application by sending thousands of HTTP
// requests randomly and concurrently while calling Go Mock Yourself Mocks management functions
//

func TestGoMockYourselfConcurrency(t *testing.T) {
	//
	// This test is skipped as Go Mock Yourself is currently not thread safe, kindly
	// read related documentation notes so you understand how to use it safely
	//

	t.Skip()

	//
	// This tests just asserts there are no panics with tons of forced concurrency collision attempts
	//

	defer func() {
		if panic := recover(); panic != nil {
			t.Errorf("Go Mock Yourself seems not to be thread safe! Come on fatty, put your shit together!")
		}
	}()

	//
	// Initialize Channel used by runRandomly() functions to know when to stop all its running functions
	//

	stopRunRandomlyFunctions := make(chan bool)

	//
	// Run randomly Go Mock Yourself management functions while tons of requests are being received
	//

	go runRandomly(stopRunRandomlyFunctions, time.Microsecond * 3, func() {
		go_mock_yourself_http.Reset()
	})

	go runRandomly(stopRunRandomlyFunctions, time.Microsecond * 1, func() {
		mock := e2e_helpers.CreateMock(nil, ".*", nil, nil)

		//
		// Randomly remove Mock Response making Go Mock Yourself redirect request to Go's http package Do() function
		//

		if go_mock_yourself_helpers.RandomInt(1, 20) % 3 == 0 {
			mock.Response = nil
		}

		go_mock_yourself_http.Install(mock)
	})

	//
	// Sync "Waiter" Pattern initialization so test case waits for all concurrent HTTP requests to finish
	//

	const totalIssuesOfTonsOfRequests = 10

	var threadsWaiter sync.WaitGroup
	threadsWaiter.Add(totalIssuesOfTonsOfRequests)

	//
	// Issue tons of requests and just ensure
	//

	for i := 0; i < totalIssuesOfTonsOfRequests; i++ {
		go issueTonsOfRequests(&threadsWaiter)
		time.Sleep(time.Millisecond * time.Duration(35))
	}

	//
	// Wait all HTTP Requests are finished
	//

	threadsWaiter.Wait()

	//
	// Stop runRandomly() running functions
	//

	close(stopRunRandomlyFunctions)

	//
	// Just wait a couple of seconds to ensure runRandomly() functions finished
	//

	time.Sleep(time.Second * 5)

	//
	// Bad anarchist, i mean, good citizen
	//

	go_mock_yourself_http.Reset()
}

//
// issueTonsOfRequests() will issue randomly tons of HTTP requests while meanwhile attempting
// to alter Go Mock Yourself internal shared resources ensuring application is thread safe
//

func issueTonsOfRequests(threadsWaiter *sync.WaitGroup) {
	//
	// Flag issueTonsOfRequests() finished
	//

	defer threadsWaiter.Done()

	//
	// Initialise Channel that will track each HTTP request end
	//

	httpRequestDone := make(chan bool)

	//
	// Issue HTTP Requests
	//

	const totalHTTPRequests = 10

	for i := 0; i < totalHTTPRequests; i++ {
		for _, method := range e2e_helpers.SupportedHTTPMethods {
			go func(method string) {
				httpRequestError := fmt.Errorf("almighty god of loops kindly let me in")

				//
				// System opened file descriptors limit usually avoids total requests being done,
				// that's why you see this ugly loop reattempting requests until no error is returned
				//

				for httpRequestError != nil {
					//
					// Sleep random nanoseconds (we want all requests mixed up and being run concurrently)
					//

					randomInt := go_mock_yourself_helpers.RandomInt(15, 50)
					time.Sleep(time.Millisecond * time.Duration(randomInt))

					//
					// Generate HTTP Request
					//

					request := e2e_helpers.CreateHTTPRequest(method, "/my-bogus-url", nil, "")

					//
					// Send Request
					//

					_, httpRequestError = http.DefaultClient.Do(request)
				}

				//
				// HTTP Request finished :)
				//

				httpRequestDone <- true

			}(method)
		}
	}

	//
	// Wait until all issued HTTP requests finished
	//

	httpRequestsDone := 0
	httpRequestsTotal := totalHTTPRequests * len(e2e_helpers.SupportedHTTPMethods)

	for httpRequestsDone != httpRequestsTotal {
		select {
		case <-httpRequestDone:
			httpRequestsDone++
		}
	}

	close(httpRequestDone)
}

//
// runRandomly() will run the passed function every X random microseconds until signaled to stop through stopRandomRun
//

func runRandomly(stopRandomRun chan bool, interval time.Duration, fn func()) {
	stopRandomRunClosed := false
	ticker := time.NewTicker(interval)

	for !stopRandomRunClosed {
		select {
		case <- stopRandomRun:
			stopRandomRunClosed = true

		case <- ticker.C:
			fn()
		}
	}

	ticker.Stop()
}
