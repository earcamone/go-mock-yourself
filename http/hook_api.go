package go_mock_yourself_http

//
// Install Mock
//
//
func Install(mock *Mock) error {
	return mockingSchemeMocks.Add(mock)
}

//
// Reset Go Mock Yourself's Mocking Scheme Mocks list
//
//
func Reset() {
	mockingSchemeMocks.Reset()
}

//
// Play() will enable Go Mock Yourself's HTTP mocking scheme if paused using the Pause() function
//
//
func Play() {
	mockingSchemePaused.SetTo(false)
}

//
// Pause() will pause Go Mock Yourself's HTTP mocking scheme, meaning it will stop working until
// its enabled again using the Play() function
//
//
func Pause() {
	mockingSchemePaused.SetTo(true)
}
