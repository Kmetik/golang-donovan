package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok {
		fmt.Printf("%T is type of b and %T is type of w\n", f, w)

	}
	if b, ok := w.(*bytes.Buffer); !ok {
		fmt.Printf("%T is type of b and %T is type of w", b, w)
	}
}
