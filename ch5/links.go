package ch5

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" && a.Val != "#" {
				links = append(links, a.Val)
			}
		}
	}
	if node.FirstChild != nil {
		links = visit(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		links = visit(links, node.NextSibling)
	}
	return links
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("response is dead, %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server %s responded with status %d", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("not able to parse response body: %s", err)
	}

	var links []string
	visitNode := func(node *html.Node) bool {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
		return true
	}
	ForEachNode(doc, visitNode, nil)
	defer resp.Body.Close()
	return links, nil
}
