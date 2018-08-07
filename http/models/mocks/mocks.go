package go_mock_yourself_mocks

import (
	"fmt"
	"sync"
)

//
// Mocks Thead-Safe Slice
//

type Mocks struct {
	mocks         []Mock
	mocksMutex    sync.Mutex
}

//
// Add() will insert a Mock into the Mocks Slice
//

func (self *Mocks) Add(mock *Mock) error {
	if mock == nil {
		return fmt.Errorf("Can't add nil Mock")
	}

	//
	// Was mock initialized properly?
	//

	if mockError := mock.Ready(); mockError != nil {
		return fmt.Errorf("Can't add Mock, not initialized properly: %s", mockError.Error())
	}

	//
	// Install Mock!
	//

	self.mocksMutex.Lock()
	self.mocks = append(self.mocks, *mock)
	self.mocksMutex.Unlock()

	return nil
}

//
// Reset() will remove all Mocks from the Mocks Slice
//

func (self *Mocks) Reset() {
	self.mocksMutex.Lock()
	self.mocks = []Mock{}
	self.mocksMutex.Unlock()
}

//
// range operator thread-safe Mocks Iterator
//

func (self *Mocks) Iterator() <-chan Mock {
	self.mocksMutex.Lock()
	defer self.mocksMutex.Unlock()

	sliceChannel := make(chan Mock, len(self.mocks))

	for _, mock := range self.mocks {
		sliceChannel <- mock
	}

	close(sliceChannel)
	return sliceChannel
}
