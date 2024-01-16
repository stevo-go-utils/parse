# parse
A simple way to parse HTML string to queryable nodes to quickly navigate to specific nodes and their values in Go.

# How to use
To convert HTML string to an interactable object ,`*parse.Node`, use the `Parse()` function. This function will only error if the HTML input is invalid HTML.
```go
node, err := parse.Parse(htmlInput)
if err != nil {
    panic(err)
}
```

# Queries
To search of child nodes of the root this package uses css selector queries supported by the [cascadia package](https://github.com/andybalholm/cascadia).

## Examples
#### *parse.Node.Query(query string) (res *Node, err error)
Takes in css selector as a string, and queries the root node for any matching nodes. Returns the first matching node. Returns an error if the function was unable to find a valid node or the css selector was invalid.
```go
res, err := node.Query(`div[id="foo"]`)
if err != nil {
    panic(err)
}
fmt.Println(res.Render())
```
#### *parse.Node.QueryAll(query string) (res []*Node, err error)
Takes in css selector as a string, and queries the root node for any matching nodes. Returns all matching nodes in an array. Returns an error if the css selector was invalid.
```go
resArr, err := node.QueryAll(`div[class="foo"]`)
if err != nil {
    panic(err)
}
for _, res := range resArr {
    fmt.Println(res.Render())
}
```

# Utility Methods

#### *parse.Node.InnerHtml() (innerHtml string)
Grabs the first child text node and its corresponding data.
```go
html := `<div>Foo</div>`
node, err := parse.Parse(htmlInput, true)
if err != nil {
    panic(err)
}
fmt.Println(node.InnerHtml())
```

