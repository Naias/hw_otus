package hw02_unpack_string //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

/*
Test checks if multibyte chars are handled properly: dataset contains russian, japanese and unicode pict symbols
 */
func TestMultibyteChars(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "кот",
			expected: "кот",
		},
		{
			input:    "я2на",
			expected: "яяна",
		},
		{
			input:    "平仮名3",
			expected: "平仮名名名",
		},
		{
			input:    "平13",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "★5★0",
			expected: "★★★★★",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

/*
Test checks if capital letters are handled properly
*/
func TestCapitalChars(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "AA",
			expected: "AA",
		},
		{
			input:    "rhZ9krya",
			expected: "rhZZZZZZZZZkrya",
		},
		{
			input:    "УтКи2",
			expected: "УтКии",
		},
		{
			input:    "★5★0D1uu2ck",
			expected: "★★★★★Duuuck",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

/*
Test checks if whitespaces are handled properly
*/
func TestEmpty(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    " ",
			expected: " ",
		},
		{
			input:    "a4b c2d5e",
			expected: "aaaab ccddddde",
		},
		{
			input:    "平 1",
			expected: "平 ",
		},
		{
			input:    "a4bc 2d5e",
			expected: "aaaabc  ddddde",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackWithEscape(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
