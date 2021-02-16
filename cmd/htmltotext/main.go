package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"entf.net/htmltools"
)

func main() {
	htmltools.Main(os.Args[1:], func(doc *html.Node) {
		for n := range htmltools.FindRecursive(
			doc,
			htmltools.MatchNodeTypeFunc(html.TextNode)) {
			if t := strings.TrimSpace(n.Data); t != "" {
				fmt.Println(t)
			}
		}
	})
}
