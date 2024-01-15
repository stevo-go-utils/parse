package parse_test

import (
	"regexp"
	"testing"

	"github.com/stevo-go-utils/parse"
)

func TestAttrsVal(t *testing.T) {
	// find first <a> tag element with href attribute and return the href value
	val, err := parse.AttrsVal(htmlInput, []string{"href", "class"}, parse.TagNameAttrValOpt("a"))
	if err != nil {
		t.Fatal(err)
	}
	if val["href"] != "https://bulbagarden.net" || val["class"] != "bg-global-nav-logo" {
		t.Fatalf("expected map[class:bg-global-nav-logo href:https://bulbagarden.net], got '%s'", val)
	}
}

func TestAttrsVals(t *testing.T) {
	// find all <a> tag elements with href attribute and return their href value
	vals := parse.AttrsVals(htmlInput, []string{"href", "class"}, parse.TagNameAttrValOpt("a"))
	if len(vals) != 245 {
		t.Fatalf("expected 981 urls, got %d", len(vals))
	}
}

func TestAttrsValOpts(t *testing.T) {
	// add in test cases here for the various options
	pattern, err := regexp.Compile(`^Bulba*`)
	if err != nil {
		t.Fatal(err)
	}
	vals := parse.AttrVals(htmlInput, "href", parse.TagNameAttrValOpt("a"), parse.InnerHtmlRegexAttrValOpt(pattern))
	if len(vals) != 12 {
		t.Fatalf("expected 12 urls, got %d", len(vals))
	}
}
