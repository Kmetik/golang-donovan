package ch5

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var file *os.File

func init() {
	opened, err := os.Open("chlen.html")
	if err != nil {
		os.Exit(0)
	}
	file = opened
}

func FindLinkMain() {
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	findlinks1(doc)
	defer file.Close()
}

func FindLinks2Main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks2(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findlinks1(doc *html.Node) {
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func findlinks2(url string) ([]string, error) {
	err := WaitForServer(url)
	if err != nil {
		log.Fatalf("vseyo sodhlo shef %s", err)
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("shef vseyo propalo %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("shef vseyo propalo %s", err)

	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("shef vseyo propalo %s", err)
	}
	defer resp.Body.Close()
	return visit(nil, doc), nil
}

// 5.2
func CountNodes() {
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	counterMap := make(map[string]int)
	nodesCounter(counterMap, doc)
	defer file.Close()
	for k, v := range counterMap {
		fmt.Println(k, v)
	}
}

func nodesCounter(counter map[string]int, node *html.Node) {
	if node.Type == html.ElementNode {
		counter[node.Data]++
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodesCounter(counter, child)
	}
}

//	func Expand(s string, replacer func(value string) string) string {
//		s = strings.Replace(s, "$foo", replacer("hui"), -1)
//		return s
//	}
//
// / какое-то говно и не понимаю задачу
func Expand(s string, f func(string) string) string {
	items := strings.Split(s, " ")
	for i, item := range items {
		if strings.HasPrefix(item, "$") {
			items[i] = f(item[1:])
		}
	}
	return strings.Join(items, " ")
}

func Outline(stack []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
	}

	if node.NextSibling != nil {
		stack = Outline(stack, node.NextSibling)
	}

	if node.FirstChild != nil {
		stack = Outline(stack, node.FirstChild)
	}
	return stack
}

func ForEachNode(node *html.Node, pre, post func(node *html.Node) bool) {
	if pre != nil {
		if ok := pre(node); !ok {
			return
		}

	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ForEachNode(child, pre, post)
	}
	if post != nil {
		if ok := post(node); !ok {
			return
		}
	}
}

func ElementById(node *html.Node, id string) *html.Node {
	for _, attr := range node.Attr {
		if attr.Key == "id" && attr.Val == id {
			return node
		}
	}
	return nil
}
