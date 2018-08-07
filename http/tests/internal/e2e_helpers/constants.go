package e2e_helpers

import (
	"net/http"
)

//
// Supported HTTP Methods
//

var SupportedHTTPMethods = []string {
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}
