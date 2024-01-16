package parse

import (
	"bytes"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

type Node struct {
	*html.Node
	InnerHtml string
	Children  []*Node
}

func Parse(body string) (node *Node, err error) {
	var n *html.Node
	n, err = html.Parse(strings.NewReader(body))
	if err != nil {
		return
	}
	node = &Node{Node: n}
	node.load()
	return
}

func (node *Node) load() {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		n := &Node{Node: child}
		n.load()
		if child.Type == html.TextNode {
			n.InnerHtml = child.Data
		}
		node.Children = append(node.Children, n)
	}
}

func (node *Node) Query(query string) *Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return node
	}
	cascadia.Query(node.Node, sel)
}

func (node *Node) Render() string {
	var buf bytes.Buffer
	html.Render(&buf, node.Node)
	return buf.String()
}
