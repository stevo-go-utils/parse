package parse

import (
	"bytes"
	"errors"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

type Node struct {
	*html.Node
}

func Parse(body string) (node *Node, err error) {
	var n *html.Node
	n, err = html.Parse(strings.NewReader(body))
	if err != nil {
		return
	}
	node = &Node{Node: n}
	return
}

func (node *Node) Query(query string) (res *Node, err error) {
	sel, err := cascadia.Parse(query)
	if err != nil {
		err = errors.New("failed to parse query")
		return
	}
	htmlRes := cascadia.Query(node.Node, sel)
	if htmlRes == nil {
		err = errors.New("no results found")
	}
	res = &Node{Node: htmlRes}
	return
}

func (node *Node) Render() string {
	var buf bytes.Buffer
	html.Render(&buf, node.Node)
	return buf.String()
}

func (node *Node) InnerHtml() (innerHtml string) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			innerHtml = c.Data
		}
	}
	return
}
