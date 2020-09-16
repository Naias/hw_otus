package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var b strings.Builder
	var prevChar rune
	// convert string to a slice of runes to properly handle multibyte chars
	input := []rune(s)
	for key, i := range input {
		if key > 0 {
			prevChar = input[key-1]
		}

		// linter is bs-ing with "deeply nested (complexity: 5)" here
		if unicode.IsDigit(i) { // nolint:nestif
			if key == 0 {
				return "", ErrInvalidString
			}

			if unicode.IsDigit(prevChar) {
				return "", ErrInvalidString
			}

			if string(i) == "0" {
				r := b.String()
				r = strings.TrimSuffix(r, string(prevChar))
				// dirty hack to keep using string builder
				b.Reset()
				b.WriteString(r)
				continue
			}

			numOfRepeats, err := strconv.Atoi(string(i))
			if err != nil {
				log.Fatalf("Error converting char %c to int: %v", i, err)
			}

			addendum := strings.Repeat(string(prevChar), numOfRepeats-1)
			b.WriteString(addendum)
		} else {
			b.WriteRune(i)
		}
	}

	result := b.String()
	return result, nil
}
