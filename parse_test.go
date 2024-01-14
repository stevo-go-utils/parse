package parse_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stevo-go-utils/parse"
)

var htmlInput string

func init() {
	data, err := os.ReadFile("./axew.html")
	if err != nil {
		fmt.Println("failed to load html input form 'axew.html'")
		panic(err)
	}
	htmlInput = string(data)
}

func TestAttrVal(t *testing.T) {
	// find first <a> tag element with href attribute and return the href value
	val, err := parse.AttrVal(htmlInput, "href", parse.AttrValTagNameOpt("a"))
	if err != nil {
		t.Fatal(err)
	}
	if val != "https://bulbagarden.net" {
		t.Fatalf("expected 'https://bulbagarden.net', got '%s'", val)
	}
}

func TestAttrVals(t *testing.T) {
	// find all <a> tag elements with href attribute and return their href value
	vals := parse.AttrVals(htmlInput, "href", parse.AttrValTagNameOpt("a"))
	if len(vals) != 981 {
		t.Fatalf("expected 981 urls, got %d", len(vals))
	}
}

func TestAttrValOpts(t *testing.T) {
	// add in test cases here for the various options
}
