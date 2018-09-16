package go_mock_yourself_helpers

import (
	"time"
	"strconv"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//
// DuplicateMap() will simply make a copy of the received map
//

func DuplicateMap(sourceMap map[string]string) map[string]string {
	mapCopy := make(map[string]string)

	for key, value := range sourceMap {
		mapCopy[key] = value
	}

	return mapCopy
}

//
// RandomString() will return a random string of the specified length
//

func RandomString(length int) string {
	randomString := ""

	for len(randomString) < length {
		randomString += strconv.Itoa(rand.Int())
	}

	return randomString[:length]
}

//
// RandomInt() will, mmm, i will let you guess this one
//

func RandomInt(min, max int) int {
	return rand.Intn(max - min) + min
}
