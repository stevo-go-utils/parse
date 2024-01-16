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

// Creates a Node from a node from the built-in html package
func NewNode(htmlNode *html.Node) *Node {
	return &Node{Node: htmlNode}
}

// Enter an html valid body and convert it to Node tree for further functionality.
func Parse(body string) (node *Node, err error) {
	var n *html.Node
	n, err = html.Parse(strings.NewReader(body))
	if err != nil {
		return
	}
	node = &Node{Node: n}
	return
}

// Grabs the body node and resorts to itself if body wasn't located
func (node *Node) Body() (res *Node) {
	var err error
	res, err = node.Query("body")
	if err != nil || res == nil || res.FirstChild == nil {
		res = node
	}
	res = NewNode(res.FirstChild)
	return
}

// Render the node and its children as html
func (node *Node) Render() string {
	var buf bytes.Buffer
	html.Render(&buf, node.Node)
	return buf.String()
}

// Render the node its siblings and its children as html
func (node *Node) RenderWithSiblings() (res string) {
	var buf bytes.Buffer
	html.Render(&buf, node.Node)
	s := node.NextSibling
	for s != nil {
		html.Render(&buf, s)
		s = s.NextSibling
	}
	return buf.String()
}

// Given a css selector query to find the first node that matches the query. Errors if fails to find one or invalid css selector query.
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

// Given a css selector query to finds all children nodes that match the query. Errors if invalid css selector query.
func (node *Node) QueryAll(query string) (res []*Node, err error) {
	sel, err := cascadia.Parse(query)
	if err != nil {
		err = errors.New("failed to parse query")
		return
	}
	htmlRes := cascadia.QueryAll(node.Node, sel)
	for _, htmlNode := range htmlRes {
		res = append(res, &Node{Node: htmlNode})
	}
	return
}

// Return the innerHtml of the node, by finding the first direct TextNode child of the Node. The search will not try and find innerHtml of the node's children.
func (node *Node) InnerHtml() (innerHtml string) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			innerHtml = c.Data
		}
	}
	return
}

// Given an attribute key (e.g. `class`) returns whether the node contains that attribute and the value of such attribute
func (node *Node) GetAttr(key string) (val string, has bool) {
	for _, attr := range node.Attr {
		if attr.Key == key {
			has = true
			val = attr.Val
			return
		}
	}
	return
}

// Given an attribute key (e.g. `class`) returns value of such attribute and defaults to an empty string
func (node *Node) MustGetAttr(key string) (val string) {
	for _, attr := range node.Attr {
		if attr.Key == key {
			val = attr.Val
			return
		}
	}
	return
}
