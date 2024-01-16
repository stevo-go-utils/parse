package parse_test

import (
	"fmt"
	"os"
)

var (
	htmlInput  string
	htmlInput2 string = `<div>Foo</div><p>Bar</p>`
)

func init() {
	data, err := os.ReadFile("./axew.html")
	if err != nil {
		fmt.Println("failed to load html input form 'axew.html'")
		panic(err)
	}
	htmlInput = string(data)
}
