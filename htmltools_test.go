package htmltools

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func getdoc(source string) *html.Node {
	doc, err := html.Parse(strings.NewReader(source))
	if err != nil || doc == nil || doc.FirstChild == nil {
		panic(err)
	}
	return doc
}

func TestBody(t *testing.T) {
	body, err := Body(getdoc(`<!DOCTYPE html><html>
<head><title>test</title></head>
<body>SUCCESS</body></html>`))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if body == nil {
		t.Fatal("body is nil")
	}
	if body.Type != html.ElementNode {
		t.Fatal("body is not an element")
	}
	if body.FirstChild == nil {
		t.Fatal("body has no children")
	}
	if body.FirstChild.Data != "SUCCESS" {
		t.Fatal("could not find body text")
	}

	body, err = Body(body)
	if err != ErrNodeIsNotADocumentNode {
		t.Fatal("no error on invalid node")
	}
	if body != nil {
		t.Fatal("body not nil on invalid node")
	}
}

func TestFindRecursive(t *testing.T) {
	doc := getdoc(`<!DOCTYPE html>
<div><div><div><p></p></div><p></p></div><p></p></div><p></p>`)
	ch := FindRecursive(doc, func(n *html.Node) bool {
		return n.Data == "p"
	})
	c := 0
	for _ = range ch {
		c++
	}
	if c != 4 {
		t.Fatalf("4 nodes expected, found %d", c)
	}

	doc = getdoc(`<!DOCTYPE html><html>
<div><div><div><p>Hello</p></div><p></p></div><p></p></div><p>World</p>`)
	ch = FindRecursive(doc, MatchNodeTypeFunc(html.TextNode))
	c = 0
	for _ = range ch {
		c++
	}
	if c != 2 {
		t.Fatalf("2 nodes expected, found %d", c)
	}
}

func TestAttr(t *testing.T) {
	body, err := Body(getdoc(`<!DOCTYPE html>
<div id="1"></div>
<div id="1"></div>
TEST
<div ID="1"></div>
<div><div id="0"></div></div>`))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if body == nil {
		t.Fatal("body is nil")
	}
	topNodes := Children(body)
	ids, err := Attr([]string{"id"}, topNodes...)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(ids) != 3 {
		t.Fatalf("3 attributes expected, found %d", len(ids))
	}
	for _, set := range ids {
		if len(set) != 1 {
			t.Fatalf("1 attribute in set expected, found %d", len(set))
		}
		if set[0] != "1" {
			t.Fatal("invalid attribute value")
		}
	}

	body, err = Body(getdoc(`<!DOCTYPE html>
<div id="1" class="test"></div>
<div id="1" data-test="data"></div>
<div ID="1"></div>`))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if body == nil {
		t.Fatal("body is nil")
	}
	topNodes = Children(body)
	ids, err = Attr([]string{"id", "data-test", "class"}, topNodes...)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(ids) != 3 {
		t.Fatalf("3 attributes expected, found %d", len(ids))
	}
	for _, set := range ids {
		if len(set) != 3 {
			t.Fatalf("3 attribute in set expected, found %d", len(set))
		}
		if set[0] != "1" {
			t.Fatal("invalid attribute value")
		}
		if set[1] != "" && set[1] != "data" {
			t.Fatal("invalid attribute value")
		}
		if set[2] != "" && set[2] != "test" {
			t.Fatal("invalid attribute value")
		}
	}
}

func TestIndentHeadings(t *testing.T) {
	body, err := Body(getdoc(`<!DOCTYPE html>
<h1></h1>TEST<div></div><h2></h2><h6></h6>`))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if body == nil {
		t.Fatal("body is nil")
	}
	topNodes := Children(body)
	err = IndentHeadings(3, topNodes...)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for _, n := range topNodes {
		if n.Type != html.ElementNode {
			continue
		}
		switch n.Data {
		case "div", "h4", "h5", "h6":
		default:
			t.Fatalf("invalid node %s", n.Data)
		}
	}

	err = IndentHeadings(-5, topNodes...)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for _, n := range topNodes {
		if n.Type != html.ElementNode {
			continue
		}
		switch n.Data {
		case "div", "h1", "h2":
		default:
			t.Fatalf("invalid node %s", n.Data)
		}
	}
}

func TestUnwrap(t *testing.T) {
	body, err := Body(getdoc(`<!DOCTYPE html>
<div>TEST<!--TEST--></div>`))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if body == nil {
		t.Fatal("body is nil")
	}
	err = Unwrap(body.FirstChild)
	if err != nil {
		t.Fatalf("%v", err)
	}
	topNodes := Children(body)
	if len(topNodes) != 2 {
		t.Fatalf("2 nodes expected, found %d", len(topNodes))
	}
}
