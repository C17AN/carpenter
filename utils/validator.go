package utils

import (
	"errors"
	"unicode"
)

// [Valid image tag rules are described here]
// https://docs.docker.com/engine/reference/commandline/tag/#extended-description

func isASCII(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func ImageTagValidator(input string) error {
	if isASCII(input) == false {
		return errors.New("Tag must consist of ASCII char")
	} else if len(input) > 128 {
		return errors.New("Tag length should be shorter than 128byte")
	} else if unicode.IsSymbol(rune(input[0])) {
		return errors.New("Tag cannot be started with symbolic char")
	}
	return nil
}
