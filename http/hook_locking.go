package go_mock_yourself_http

import (
	"sync"

	// Third-Party Imports
	"github.com/bouk/monkey"
)

//
// Mocking Scheme's Go http package sole hook
//
// NOTE: "under the hood", all http package helper functions (http.Get, http.Post) call an http client's
// Do() method, hook PatchGuard holds the hook for all available http client instances Do() methods
//

var hook *monkey.PatchGuard

//
// Hook Mutex (Monkey Patching library states its not thread-safe)
//

var hookMutex sync.Mutex

//
// hookInstall() will restore Go's native HTTP package hook thread-safe
//

func hookInstall() {
	hook.Restore()
	hookMutex.Unlock()
}

//
// hookUninstall() will restore Go's native HTTP requests management scheme thread-safe
//

func hookUninstall() {
	hookMutex.Lock()
	hook.Unpatch()
}
