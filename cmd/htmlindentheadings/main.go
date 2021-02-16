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
		for node := range htmltools.FindRecursive(doc, nil) {
			htmltools.IndentHeadings(lvls, node)
		}
		html.Render(os.Stdout, doc)
	})
}
