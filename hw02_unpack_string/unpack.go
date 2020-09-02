package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

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
				continue
			}

			numOfRepeats, err := strconv.Atoi(string(i))
			if err != nil {
				log.Fatalf("Error converting char %c to int: %v", i, err)
			}

			result += strings.Repeat(string(prevChar), numOfRepeats-1)
		} else {
			//fmt.Printf("\n%v:%c is string, prev: %c", key, i, prevChar)

			result += string(i)
		}
	}
	// Place your code here
	return result, nil
}
