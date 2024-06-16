package ch2

import (
	"flag"
	"fmt"
	"strings"
)

var b = flag.Bool("b", false, "boolean value")
var s = flag.String("s", " ", "string value")

func Flags() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), " "))
	if !*b {
		fmt.Println(*b)
	}
	fmt.Println(*s, *b)
}
