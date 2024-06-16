package ch4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// собирает повторы строк в стд.вводе
func Dedup() {
	in := bufio.NewScanner(os.Stdin)
	seen := make(map[string]bool)
	for in.Scan() {
		line := in.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := in.Err(); err != nil {
		os.Exit(1)
	}
}

func Charcount(path string) {
	count := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	var invalid int
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occured %s\n %s\n", err, path)
		os.Exit(1)
	}
	buf := bufio.NewReader(file)
	for {
		r, n, err := buf.ReadRune()
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		count[r]++
		utflen[n]++
		if err == io.EOF {
			fmt.Println("file end")
			break
		}
	}

	for kc, vc := range count {
		fmt.Printf("%q: %d\n", kc, vc)
	}

	for iu, vu := range utflen {
		fmt.Printf("%d: %d\n", iu, vu)
	}
	fmt.Println(invalid)
}

var graph = make(map[string]map[string]bool)

func AddEdge(from string, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func HasEdge(from string, to string) bool {
	return graph[from][to]
}

func CharcountByType(path string) {
	counts := make(map[string]int)
	if file, err := os.Open(path); err == nil {
		buf := bufio.NewReader(file)

		for {
			r, _, err := buf.ReadRune()
			if !utf8.ValidRune(r) {
				continue
			}
			if unicode.IsLetter(r) {
				counts["letters"]++
			}
			if unicode.IsDigit(r) {
				counts["numbers"]++
			}
			if err == io.EOF {
				break
			}
		}

		for k, v := range counts {
			fmt.Printf("%q: %d\n", k, v)
		}
	} else {
		fmt.Fprintf(os.Stderr, "error occured with %s: %s\n", path, err)
		os.Exit(1)
	}
}

func WordFreq(path string) {
	counts := make(map[string]int)

	if file, err := os.Open(path); err == nil {
		var in = bufio.NewScanner(file)
		in.Split(bufio.ScanWords)

		for in.Scan() {
			word := in.Text()
			counts[word]++
			// if _, ok := counts[word]; !ok {
			// 	counts[word]++
			// }
		}
		fmt.Println(counts)
	} else {
		os.Exit(1)
	}
}
