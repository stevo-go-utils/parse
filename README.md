# parse
A simple way to parse HTML string to queryable nodes to quickly navigate to specific nodes and their values in Go.

# How to install
`go get github.com/stevo-go-utils/parse`

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

#### *parse.Node.Render() (res string)
Renders the node and its children.
```go
html := `<div>Foo</div>`
node, err := parse.Parse(html)
if err != nil {
    panic(err)
}
fmt.Println(node.Body().Render()) // <div>Foo</div>
```

#### *parse.Node.RenderWithChildren() (res string)
Renders the node and its children. Then appends the Rendered siblings of the root node if any exist.
```go
html := `<div>Foo</div><p>Bar</p>`
node, err := parse.Parse(html)
if err != nil {
    panic(err)
}
fmt.Println(node.Body().RenderWithChildren()) // <div>Foo</div><p>Bar</p>
```

#### *parse.Node.InnerHtml() (innerHtml string)
Grabs the first child text node and its corresponding data.
```go
html := `<div>Foo</div>`
node, err := parse.Parse(html)
if err != nil {
    panic(err)
}
fmt.Println(node.Body().InnerHtml()) // Foo
```

#### *parse.Node.GetAttr(key string) (val string, has bool)
Given an attribute key, returns whether the node contains that attribute and the value of such attribute.
```go
html := `<div id="foo"></div>`
node, err := parse.Parse(html)
if err != nil {
    panic(err)
}
val, has := node.GetAttr("id")
if !has {
    panic(errors.New("failed to find attr value"))
}
fmt.Println(val) // foo
```

#### *parse.Node.MustGetAttr(key string) (val string, has bool)
Given an attribute key, returns value of such attribute and defaults to an empty string.
```go
html := `<div id="foo"></div>`
node, err := parse.Parse(html)
if err != nil {
    panic(err)
}
val := node.MustGetAttr("id")
fmt.Println(val) // foo
val = node.MustGetAttr("class")
fmt.Println(val) // EMPTY STRING
```


