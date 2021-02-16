package main // import "entf.net/htmltools/htmlattr"

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"entf.net/htmltools/shared"
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
	for i, attr := range attrs {
		attrs[i] = strings.ToLower(attr)
	}
	shared.Main(args[1:], func(doc *html.Node) {
		var body *html.Node
		for n := doc.FirstChild.FirstChild; n != nil; n = n.NextSibling {
			if strings.ToLower(n.Data) == "body" {
				body = n
				break
			}
		}
		if body == nil {
			fmt.Fprintln(os.Stderr, "document does not contain a body")
			os.Exit(1)
		}
		for n := body.FirstChild; n != nil; n = n.NextSibling {
			if n.Type != html.ElementNode {
				continue
			}
			list := make([]string, len(attrs))
			var any bool
			for i, attrn := range attrs {
				for _, attr := range n.Attr {
					if strings.ToLower(attr.Key) == attrn {
						any = true
						list[i] = attr.Val
					}
				}
			}
			line := strings.Join(list, fs)
			if any {
				fmt.Println(line)
			}
		}
	})
}
