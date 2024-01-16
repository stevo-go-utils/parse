package parse_test

import (
	"testing"

	"github.com/stevo-go-utils/parse"
)

func TestBody(t *testing.T) {
	body, err := parse.Parse(htmlInput)
	if err != nil {
		t.Fatal(err)
	}
	_ = body
}
