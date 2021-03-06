package main

import (
	"fmt"
	"os"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"

	"entf.net/htmltools"
	"entf.net/htmltools/cmd"
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
	cmd.Main(args[1:], func(doc *html.Node) {
		unwrap(sel, doc)
	})
}

func unwrap(sel cascadia.Selector, doc *html.Node) {
	for _, n := range sel.MatchAll(doc) {
		if err := htmltools.Unwrap(n); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
	html.Render(os.Stdout, doc)
}
