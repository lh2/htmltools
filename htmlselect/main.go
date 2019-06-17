package main // import "entf.net/htmltools/htmlselect"

import (
	"fmt"
	"os"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"

	"entf.net/htmltools/shared"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: htmlselect SELECTOR [FILES...]")
		os.Exit(1)
	}
	sel, err := cascadia.Compile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "selector invalid: %v\n", err)
		os.Exit(1)
	}
	shared.Main(args[1:], func(doc *html.Node) {
		dosel(sel, doc)
	})
}

func dosel(sel cascadia.Selector, doc *html.Node) {
	for _, n := range sel.MatchAll(doc) {
		html.Render(os.Stdout, n)
	}
}
