package main

import (
	"fmt"
	"os"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"

	"entf.net/htmltools"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: htmlremove SELECTOR [FILES...]")
		os.Exit(1)
	}
	sel, err := cascadia.Compile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "selector invalid: %v\n", err)
		os.Exit(1)
	}
	htmltools.Main(args[1:], func(doc *html.Node) {
		unwrap(sel, doc)
	})
}

func unwrap(sel cascadia.Selector, doc *html.Node) {
	for _, n := range sel.MatchAll(doc) {
		cs := make([]*html.Node, 0)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			cs = append(cs, c)
		}
		for _, c := range cs {
			n.RemoveChild(c)
			n.Parent.InsertBefore(c, n)
		}
		n.Parent.RemoveChild(n)
	}
	html.Render(os.Stdout, doc)
}
