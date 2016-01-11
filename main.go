// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for i, link := range visit(nil, doc) {
		fmt.Printf("%v, %s\n", i, link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = mapping(links, n.FirstChild)
	}
	return links
}

func mapping(links []string, n *html.Node) []string {
	links = visit(links, n)
	if n.NextSibling != nil {
		links = mapping(links, n.NextSibling)
	}
	return links
}
