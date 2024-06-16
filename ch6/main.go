package main

import (
	"fmt"
	bytecounter "study/donovan/ch6/bytcounter"
)

type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func CountAndSplitWords(data []byte, atEOF bool) (advance int, token []byte, err error) {

	return 0, nil, nil
}

func main() {
	var w bytecounter.WordsCounter
	w.Write([]byte("hello bitch suka "))
	fmt.Println(w)
}
