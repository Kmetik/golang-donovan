package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

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

var tokens = make(chan struct{}, 20)

func crawl(link string) []string {
	fmt.Println(link)
	tokens <- struct{}{}
	urls, err := Extract(link)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return urls
}

func main() {
	type linkItem struct {
		uri   string
		depth int8
	}
	worklist := make(chan []linkItem)
	unseenlinks := make(chan linkItem)
	go func() {
		var items []linkItem
		for _, uri := range os.Args[1:] {
			items = append(items, linkItem{uri, 1})
		}
		worklist <- items
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenlinks {
				if link.depth > 3 {
					fmt.Printf("depth exceeded %s, depth %d\n", link.uri, link.depth)
					<-unseenlinks
					continue
				}
				founduris := crawl(link.uri)
				var foundLinks []linkItem
				for _, uri := range founduris {
					foundLinks = append(foundLinks, linkItem{uri, link.depth + 1})
				}
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.uri] {
				seen[link.uri] = true
				unseenlinks <- link
			}
		}
	}
	close(unseenlinks)
}

// crawl gopl.io >
//  gopl.io/1 crawl >
//   gopl.io/2 > depth exceeded
