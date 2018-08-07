package go_mock_yourself_http

import (
	"net/http"
)

func getMatchingMock(request *http.Request) *Mock {
	debug(" + Analysing incoming HTTP request for matching Mock: %v\n", request)

	//
	// Is Mocking Scheme temporarily paused?
	//

	if mockingSchemePaused.IsSet() {
		debug("    X Mocking Scheme Paused: Request arrival analysis skipping, redirected to Go native package\n\n")
		return nil
	}

	//
	// Does target HTTP request match any of the currently installed Mocks?
	//

	for mock := range mockingSchemeMocks.Iterator() {
		debug("    - Iterating Installed Mocks: %s\n", mock.Name)

		if mock.Request.Match(request) {
			debug("        X Iterating Installed Mocks: '%s' matches request\n\n", mock.Name)
			return &mock
		}
	}

	//
	// Target HTTP request does not match any of the installed Mocks
	//

	debug("        X Iterating Installed Mocks: no installed Mock match for request\n\n")
	return nil
}
