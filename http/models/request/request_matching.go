package go_mock_yourself_request

//
// ===================== Target Request Matching Helper methods =====================
//

import (
	"net/http"

	// Go Mock Yourself Imports
	"github.com/mercadolibre/go-mock-yourself/http/helpers"
	"github.com/mercadolibre/go-mock-yourself/http/models/internal"
)

//
// Match() will return true if the received HTTP request matches current Mock Request
//

func (self Request) Match(request *http.Request) bool {
	//
	// Retrieve HTTP request information
	//

	method, url, headers, body := go_mock_yourself_helpers.RequestInformation(request)

	//
	// NOTE: Internally, Go's http package decorates Request Headers with some default values if calling client code
	// does not specify them in its Transport information, meaning that the target server would receive this Headers
	// if client code did not specify them SO we are decorating them so current Mock's matching dumpHeaders can actually
	// be applied on a set of Headers just as they will be sent to the target server.
	//
	// TODO: maybe we should move the decorateDefaultHeaders call inside RequestInformation()
	//

	headers = go_mock_yourself_helpers.DecorateDefaultHeaders(request, headers)

	//
	// Is there a matching function for the whole request?
	//

	if self.matchingRequest != nil {
		return (*self.matchingRequest)(request)
	}

	//
	// Ensure Request matches individual Request criterias
	//

	if !self.URLMatch(url) {
		return false
	}

	if !self.MethodMatch(method) {
		return false
	}

	if !self.HeadersMatch(headers) {
		return false
	}

	if !self.BodyMatch(string(body)) {
		return false
	}

	return true
}

//
// MethodMatch() will return true if the passed method matches Mock's Request target methods
//

func (self Request) MethodMatch(method string) bool {
	match := true // if no matching criteria specified, match!

	if self.matchingMethods != nil {
		match = (*self.matchingMethods)(method)

	} else if self.methods != nil {
		match = self.methods.MatchString(method)
	}

	return match
}

//
// URLMatch() will return true if the passed url matches Mock's Request target url
//

func (self Request) URLMatch(url string) bool {
	match := true // if no matching criteria specified, match!

	if self.matchingUrl != nil {
		match = (*self.matchingUrl)(url)

	} else if self.url != nil {
		match = self.url.MatchString(url)
	}

	return match
}

//
// BodyMatch() will return true if the passed body matches Mock's Request target body
//

func (self Request) BodyMatch(body string) bool {
	match := true // if no matching criteria specified, match!

	if self.matchingBody != nil {
		match = (*self.matchingBody)(body)

	} else if self.body != nil {
		match = self.body.MatchString(body)
	}

	return match
}

//
// HeadersMatch() will return true if the passed headers match Mock's Request target headers
//

func (self Request) HeadersMatch(headers map[string]string) bool {
	match := true // if no headers regular expressions or matching function specified, force match

	if self.matchingHeaders != nil {
		match = (*self.matchingHeaders)(headers)

	} else if self.headers != nil && len(self.headers) > 0 {
		match = go_mock_yourself_models_internal_helpers.HeadersMatching(headers, self.headers)
	}

	return match
}

func (self Request) ShouldMock(request *http.Request) bool {
	return self.shouldMock == nil || self.shouldMock(request)
}
