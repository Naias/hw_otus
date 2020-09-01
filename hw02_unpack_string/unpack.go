package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

//nolint:nes
func Unpack(input string) (string, error) {
	var result string
	var prevChar rune
	for key, i := range input {
		if key > 0 {
			prevChar = rune(input[key-1])
		}

		//linter is bs-ing with "deeply nested (complexity: 5)" here
		if unicode.IsDigit(i) { // nolint:nestif
			//fmt.Printf("\n%v:%c is digit, prev: %c", key, i, prevChar)
			if key == 0 {
				return "", ErrInvalidString
			}

			if unicode.IsDigit(prevChar) {
				return "", ErrInvalidString
			}

			if string(i) == "0" {
				result = strings.TrimSuffix(result, string(prevChar))
			}

			numOfRepeats, _ := strconv.Atoi(string(i))
			if numOfRepeats > 0 {
				for j := 0; j < numOfRepeats-1; j++ {
					result += string(prevChar)
				}
			}
		} else {
			//fmt.Printf("\n%v:%c is string, prev: %c", key, i, prevChar)

			result += string(i)
		}
	}
	// Place your code here
	return result, nil
}
