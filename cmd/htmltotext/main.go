package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"entf.net/htmltools"
	"entf.net/htmltools/cmd"
)

func main() {
	cmd.Main(os.Args[1:], func(doc *html.Node) {
		for n := range htmltools.FindRecursive(
			doc,
			htmltools.MatchNodeTypeFunc(html.TextNode)) {
			if t := strings.TrimSpace(n.Data); t != "" {
				fmt.Println(t)
			}
		}
	})
}
