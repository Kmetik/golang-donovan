package ch5

import "golang.org/x/net/html"

var Iterations int16 = 0

func init() {
	Iterations = 0
}

func Min(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	min := values[0]

	for val := range values {
		if val < min {
			min = val
		}
	}

	return min
}

func MinLogN(values ...int) int {
	length := len(values)
	if length == 0 {
		return 0
	}

	if length == 2 {
		if values[0] < values[1] {
			return values[0]
		} else {
			return values[1]
		}
	}
	Iterations += 1

	subMin := MinLogN(values[2:]...)

	if values[0] < subMin {
		return values[0]
	}
	return subMin
}

func ElementsByTagName(doc *html.Node, tags ...string) []*html.Node {
	var found []*html.Node

	lookup := func(node *html.Node) bool {
		if node.Type == html.ElementNode {
			for _, tag := range tags {
				if node.Data == tag {
					found = append(found, node)
				}
			}
		}

		return true
	}

	ForEachNode(doc, lookup, nil)

	return found
}
