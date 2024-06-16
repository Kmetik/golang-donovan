package ch1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "exit" {
			break
		}
	}
	// countLines(input, counts)
	printDuplicates(counts)
}

func DupFile() {
	files := os.Args[1:]
	countsMap := make(map[string]map[string]int)
	if len(files) == 0 {
		fmt.Println("no files")
		Dup1()
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countsMap[arg] = make(map[string]int)
			countLines(f, countsMap[arg])
			f.Close()
			for key, val := range countsMap {
				fmt.Printf("Folder is %s\n", key)
				printDuplicates(val)
			}
		}
	}

}

func countLines(data *os.File, counts map[string]int) {
	input := bufio.NewScanner(data)

	for input.Scan() {
		counts[input.Text()]++
		// if input.Text() == "exit" {
		// 	break
		// }
	}
}

// memory intensive, dog shit
func DupAllFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("no files specified")
		return
	}
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		printDuplicates(counts)
	}
}

func printDuplicates(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
	}
}
