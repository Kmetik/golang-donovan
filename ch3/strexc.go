package ch3

import (
	"bytes"
)

func Comma(arg string) string {
	var buf bytes.Buffer
	n := len(arg)
	for i, r := range arg {
		left := n - i
		if left%3 != 1 {
			buf.WriteRune(r)
		} else {
			buf.WriteRune(r)
			buf.WriteRune(' ')
		}
	}
	return buf.String()
}

func Anagram(arg1 string, arg2 string) bool {
	na := len(arg1)
	nb := len(arg2)

	if na != nb {
		return false
	}

	for i := na - 1; i >= 0; i-- {
		if arg1[i] != arg2[na-1-i] {
			return false
		}
	}

	return true
}
