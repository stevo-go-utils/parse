package parse_test

import (
	"testing"

	"github.com/stevo-go-utils/parse"
	"github.com/stevo-go-utils/parse/opts/innerhtmlOpts"
)

func TestInnerHtml(t *testing.T) {
	innerHtml, err := parse.InnerHtml(htmlInput, innerhtmlOpts.InnerHtmlOpt("Bulbapedia"))
	if err != nil {
		t.Fatal(err)
	}
	if innerHtml != "Bulbapedia" {
		t.Fatalf("expected Bulbapedia, got %s", innerHtml)
	}
}
