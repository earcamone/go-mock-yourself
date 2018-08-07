package go_mock_yourself_models_internal_helpers

import (
	"fmt"
)

var IncorrectErrorParameters = []interface{} {
	3.3,
	23495829845,
	"this is not a valid integer parameter",
	[]byte("this is clearly not a valid string parameter"),
	byte(0),
}

var IncorrectRegexParameters = []interface{} {
	3.3,
	23495829845,
	fmt.Errorf("this is clearly not a valid string parameter"),
	[]byte("this is clearly not a valid string parameter"),
	byte(0),
	"(this is an invalid regex",
}

var IncorrectStringParameters = []interface{} {
	3.3,
	23495829845,
	fmt.Errorf("this is clearly not a valid string parameter"),
	[]byte("this is clearly not a valid string parameter"),
	byte(0),
}

var IncorrectIntegerParameters = []interface{} {
	"this is not a valid integer parameter",
	fmt.Errorf("this is clearly not a valid string parameter"),
	[]byte("this is clearly not a valid string parameter"),
}
