package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"entf.net/htmltools"
	"entf.net/htmltools/cmd"
	"golang.org/x/net/html"
)

func main() {
	var fs string
	flag.StringVar(&fs, "fs", ",", "field seperator")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("usage: htmlattr [-fs FIELD_SEPERATOR] ATTRIBUTES [FILES...]")
		os.Exit(1)
	}
	attrs := strings.Split(args[0], fs)
	cmd.Main(args[1:], func(doc *html.Node) {
		body, err := htmltools.Body(doc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		} else if body == nil {
			fmt.Fprintln(os.Stderr, "Document does not contain a body")
			os.Exit(1)
		}
		values, err := htmltools.Attr(attrs, htmltools.Children(body)...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		for _, v := range values {
			line := strings.Join(v, fs)
			fmt.Println(line)
		}
	})
}
