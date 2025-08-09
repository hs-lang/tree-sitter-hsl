package tree_sitter_hsl_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_hsl "github.com/hs-lang/tree-sitter-hsl/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_hsl.Language())
	if language == nil {
		t.Errorf("Error loading HyperSpace Lang grammar")
	}
}
