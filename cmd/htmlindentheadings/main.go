package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/net/html"

	"entf.net/htmltools"
)

const usage = "usage: htmlindentheadings INDENT_LEVELS [FILES...]"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}
	lvls, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(usage)
		os.Exit(1)
	}
	htmltools.Main(args[1:], func(doc *html.Node) {
		visit(lvls, doc)
		html.Render(os.Stdout, doc)
	})
}

func indent(lvls int, tag string) string {
	l := int(tag[1]) - 48
	l += lvls
	if l > 6 {
		l = 6
	}
	return fmt.Sprintf("h%d", l)
}

func visit(lvls int, n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "h1", "h2", "h3", "h4", "h5", "h6":
			n.Data = indent(lvls, n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(lvls, c)
	}
}
