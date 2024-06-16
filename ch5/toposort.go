package ch5

import (
	"fmt"
)

var Prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"linear algebra":        {"calculus"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func Toposort(graph map[string][]string) []string {
	seen := make(map[string]bool)
	var order []string
	var keys []string
	var lookup func([]string)
	lookup = func(keys []string) {
		for _, key := range keys {
			if v, ok := seen[key]; !ok {
				seen[key] = true
				lookup(graph[key])
				seen[key] = false
				order = append(order, key)
			} else if v {
				fmt.Println("loop detected")
				continue
			}
		}
	}
	for k := range graph {
		keys = append(keys, k)
	}
	lookup(keys)
	return order
}
