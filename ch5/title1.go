package ch5

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func Title(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer fmt.Println("first")
	defer fmt.Println("last")

	ct := resp.Header.Get("Content-Type")

	if ct != "text/html" && !strings.HasPrefix(ct, "text/html") {
		return fmt.Errorf("content-Type is not html")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	visitNode := func(node *html.Node) bool {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			fmt.Println(node.FirstChild.Data)
		}
		return true
	}

	ForEachNode(doc, visitNode, nil)

	return nil
}
