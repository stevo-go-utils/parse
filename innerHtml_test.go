package parse_test

import (
	"testing"

	"github.com/stevo-go-utils/parse"
)

func TestInnerHtml(t *testing.T) {
	innerHtml, err := parse.InnerHtml(htmlInput, parse.InnerHtmlInnerHtmlOpt("Bulbapedia"))
	if err != nil {
		t.Fatal(err)
	}
	if innerHtml != "Bulbapedia" {
		t.Fatalf("expected Bulbapedia, got %s", innerHtml)
	}
}
