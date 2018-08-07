package go_mock_yourself_http

import (
	"reflect"
	"net/http"

	// Third-Party Imports
	"github.com/bouk/monkey"
	"github.com/tevino/abool"
)

//
// Go Mock Yourself Mocks Slice
//

var mockingSchemeMocks *Mocks

//
// Go Mock Yourself Mocking Scheme Play/Pause bool
//
// NOTE: Meaning any HTTP call will be handled by Go's native HTTP package if mockingSchemePaused == true
//

var mockingSchemePaused *abool.AtomicBool

//
// Initialize Go Mock Yourself HTTP Mocking Scheme
//

func init() {
	//
	// Initialize "Mocking Scheme Pause" scheme's atomic bool (thread-safe)
	//

	mockingSchemePaused = abool.New()
	mockingSchemePaused.SetTo(false)

	//
	// Initialize "Mocking Scheme" Mocks Slice
	//

	mockingSchemeMocks = new(Mocks)

	//
	// Monkey patch Go http clients' Do() method which is the one used by all other package helper
	// functions (http.Get, http.Post, etc) underground to perform each specified request
	//

	var client *http.Client
	hook = monkey.PatchInstanceMethod(reflect.TypeOf(client), "Do", goMockYourselfDo)
}
