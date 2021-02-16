package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"entf.net/htmltools"
)

func main() {
	htmltools.Main(os.Args[1:], visit)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode {
		if t := strings.TrimSpace(n.Data); t != "" {
			fmt.Println(t)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}