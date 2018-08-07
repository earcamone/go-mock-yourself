package go_mock_yourself_request

import (
	"fmt"
	"regexp"
	"net/http"
)

//
// Mocking Request Information
//

type Request struct {
	//
	// Request ShouldMock() support
	//

	shouldMock func(*http.Request) bool

	//
	// Request target HTTP Methods Matching Callback (regular expressions supported)
	//

	matchingRequest *func(*http.Request) bool

	//
	// Request target HTTP methods (regular expressions supported)
	//

	methods *regexp.Regexp

	//
	// Request target HTTP Methods Matching Callback (regular expressions supported)
	//

	matchingMethods *func(string) bool

	//
	// Target Request URL (Regular Expressions supported)
	//

	url *regexp.Regexp

	//
	// Target Request URL Matching Function
	//
	// NOTE: If specified, static url regular expressions will be ignored
	//

	matchingUrl *func(string)bool

	//
	// Target Request body (Regular Expressions supported)
	//

	body *regexp.Regexp

	//
	// Request body Matching Function
	//
	// NOTE: If specified, static body regular expressions will be ignored
	//

	matchingBody *func(string) bool

	//
	// Request headers (regular expressions supported)
	//

	headers map[*regexp.Regexp]*regexp.Regexp

	//
	// Request headers Matching function
	//
	// NOTE: If specified, static headers regular expressions will be ignored
	//

	matchingHeaders *func(map[string]string) bool
}

//
// Ready() will return true if the Request is ready to be used in a Mock
//

func (self Request) Ready() error {
	requestCriteriaAvailable := self.matchingRequest != nil
	urlCriteriaAvailable := self.url != nil || self.matchingUrl != nil
	bodyCriteriaAvailable := self.body != nil || self.matchingBody != nil
	methodCriteriaAvailable := self.methods != nil || self.matchingMethods != nil
	headersCriteriaAvailable := self.headers != nil || self.matchingHeaders != nil

	//
	// At least one matching criteria should be specified in order to be ready
	//

	if !requestCriteriaAvailable && !urlCriteriaAvailable && !bodyCriteriaAvailable && !methodCriteriaAvailable && !headersCriteriaAvailable {
		return fmt.Errorf("Mock Request is not ready for usage, at least one matching criteria should be specified")
	}

	return nil
}
