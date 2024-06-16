package ch3

import (
	"fmt"
	"strings"
)

// a/b/c.go
func Basename1(arg string) {
	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] == '/' {
			arg = arg[i+1:]
			break
		}
	}

	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] == '.' {
			arg = arg[:i]
			break
		}
	}
	fmt.Println(arg)
}

func Basename2(arg string) {
	slash := strings.LastIndex(arg, "/")
	arg = arg[slash+1:]
	if dot := strings.LastIndex(arg, "."); dot >= 0 {
		arg = arg[:dot]
	}
	fmt.Println(arg)
}

func NumberRuFormat(arg string) string {
	n := len(arg)
	if n <= 3 {
		return arg
	}

	return NumberRuFormat(arg[:n-3]) + " " + arg[n-3:]
}
