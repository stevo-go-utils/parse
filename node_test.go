package parse_test

import (
	"fmt"
	"testing"

	"github.com/stevo-go-utils/parse"
)

func TestParse(t *testing.T) {
	node, err := parse.Parse(htmlInput)
	if err != nil {
		t.Error(err)
	}
	_ = node
}

func TestBody(t *testing.T) {
	node, err := parse.Parse(htmlInput2)
	if err != nil {
		t.Error(err)
	}
	if node.Body().RenderWithSiblings() != htmlInput2 {
		t.Error(fmt.Errorf("incorrect rendering"))
	}
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

func TestQueryAll(t *testing.T) {
	node, err := parse.Parse(htmlInput)
	if err != nil {
		t.Error(err)
	}
	res, err := node.QueryAll(`a[href="https://bulbagarden.net"]`)
	if err != nil {
		t.Error(err)
	}
	if len(res) != 2 {
		t.Error(fmt.Errorf("incorrect number of responses"))
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

func TestGetAttr(t *testing.T) {
	node, err := parse.Parse(htmlInput)
	if err != nil {
		t.Error(err)
	}
	res, err := node.Query(`li[class="bg-global-nav-secondary-item header"]`)
	if err != nil {
		t.Error(err)
	}
	val, has := res.GetAttr("class")
	if !has {
		t.Error(fmt.Errorf("node does not have class attribute"))
	}
	if val != "bg-global-nav-secondary-item header" {
		t.Error(fmt.Errorf("incorrect attribute value"))
	}
}

func TestMustGetAttr(t *testing.T) {
	node, err := parse.Parse(htmlInput)
	if err != nil {
		t.Error(err)
	}
	res, err := node.Query(`li[class="bg-global-nav-secondary-item header"]`)
	if err != nil {
		t.Error(err)
	}
	val := res.MustGetAttr("id")
	if val != "" {
		t.Error(fmt.Errorf("incorrect attribute value"))
	}
}
