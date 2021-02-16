package htmltools

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var (
	ErrNodeIsNotADocumentNode = errors.New("Not a document node")
	ErrNodeHasNoParent        = errors.New("Node has no parent")
)

type NodeMatchFunc func(*html.Node) bool

// Gets the body from an HTML document node.
func Body(doc *html.Node) (*html.Node, error) {
	if doc.Type != html.DocumentNode {
		return nil, ErrNodeIsNotADocumentNode
	}
	var body *html.Node
	for n := doc.FirstChild.FirstChild; n != nil; n = n.NextSibling {
		if strings.ToLower(n.Data) == "body" {
			body = n
			break
		}
	}
	return body, nil
}

// Gets all direct children.
func Children(node *html.Node) []*html.Node {
	nodes := make([]*html.Node, 0)
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		nodes = append(nodes, n)
	}
	return nodes
}

func findRecursive(node *html.Node, nodeFunc func(*html.Node) bool, ch chan<- *html.Node) {
	if nodeFunc == nil || nodeFunc(node) {
		ch <- node
	}
	for _, c := range Children(node) {
		findRecursive(c, nodeFunc, ch)
	}
}

// Returns a channel providing all nodes that match nodeFunc recursively through
// the whole document. If nodeFunc is `nil`, all nodes match.
func FindRecursive(doc *html.Node, nodeFunc NodeMatchFunc) <-chan *html.Node {
	ch := make(chan *html.Node)
	go findRecursive(doc, nodeFunc, ch)
	return ch
}

// Returns all attribite values specified in attrs for nodes.
func Attr(attrs []string, nodes ...*html.Node) ([][]string, error) {
	for i, attr := range attrs {
		attrs[i] = strings.ToLower(attr)
	}
	results := make([][]string, 0)
	for _, n := range nodes {
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
		if any {
			results = append(results, list)
		}
	}
	return results, nil
}

// Indents all headings by a certain level.
func IndentHeadings(level int, nodes ...*html.Node) error {
	for _, n := range nodes {
		switch strings.ToLower(n.Data) {
		case "h1", "h2", "h3", "h4", "h5", "h6":
		default:
			continue
		}
		l := int(n.Data[1]) - 48 //HACK: ASCII to number
		l += level
		if l > 6 {
			l = 6
		}
		n.Data = fmt.Sprintf("h%d", l)
	}
	return nil
}

// Removes node from parent and replaces it by it's children.
func Unwrap(node *html.Node) error {
	if node.Parent == nil {
		return ErrNodeHasNoParent
	}
	for _, c := range Children(node) {
		node.RemoveChild(c)
		node.Parent.InsertBefore(c, node)
	}
	node.Parent.RemoveChild(node)
	return nil
}

// Creates a NodeMatchFunc, matching a certain NodeType
func MatchNodeTypeFunc(nodeType html.NodeType) NodeMatchFunc {
	return func(node *html.Node) bool {
		return node.Type == nodeType
	}
}
