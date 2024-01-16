package parse_test

import (
	"testing"

	"github.com/stevo-go-utils/parse"
)

func TestNode(t *testing.T) {
	node, err := parse.Parse(htmlInput)
	if err != nil {
		t.Error(err)
	}
	_ = node
}

func TestQuery(t *testing.T) {
	node, err := parse.Parse(htmlInput)
	if err != nil {
		t.Error(err)
	}
	res, err := node.Query(`li[class="bg-global-nav-secondary-item header"]`)
	if err != nil {
		t.Error(err)
	}
	if res.Render() != `<li class="bg-global-nav-secondary-item header">Pokémon</li>` {
		t.Error("wrong node")
	}
}

func TestInnerHtml(t *testing.T) {
	node, err := parse.Parse(htmlInput)
	if err != nil {
		t.Error(err)
	}
	res, err := node.Query(`li[class="bg-global-nav-secondary-item header"]`)
	if err != nil {
		t.Error(err)
	}
	if res.InnerHtml() != "Pokémon" {
		t.Error("wrong inner html")
	}
}
