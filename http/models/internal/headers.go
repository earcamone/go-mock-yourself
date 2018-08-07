package go_mock_yourself_models_internal_helpers

import (
	"regexp"
)

//
// HeadersMatching() will return true if the received map[string]string
// Headers match the received map[regex]regex Headers criteria
//

func HeadersMatching(headers map[string]string, regexHeaders map[*regexp.Regexp]*regexp.Regexp) bool {
	for regexHeader, regexValue := range regexHeaders {
		matchingHeader := false

		for header, value := range headers {
			if regexHeader.MatchString(header) && regexValue.MatchString(value) {
				matchingHeader = true
				break
			}
		}

		if !matchingHeader {
			return false
		}
	}

	return true
}

//
// CompileHeaders() will compile the received Headers map[string]string to a Regular Expressions map
//

func CompileHeaders(headers map[string]string) (map[*regexp.Regexp]*regexp.Regexp, error) {
	headersRegexes := make(map[*regexp.Regexp]*regexp.Regexp)

	for header, value := range headers {
		valueRegex, regexError := regexp.Compile(value)

		if regexError != nil {
			return nil, regexError
		}

		headerRegex, regexError := regexp.Compile(header)

		if regexError != nil {
			return nil, regexError
		}

		headersRegexes[headerRegex] = valueRegex
	}

	return headersRegexes, nil
}
