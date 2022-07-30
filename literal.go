package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func basicLit(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BasicLit) {
	fmt.Fprintf(w, "%s%sBasicLit\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Kind = %s\n", parentPrefix, prefixes[1], node.Kind)
	fmt.Fprintf(w, "%s%s└── Value = %s\n", parentPrefix, prefixes[1], node.Value)
}

func compositeLit(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.CompositeLit) {
	fmt.Fprintf(w, "%s%sCompositeLit\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Type\n", parentPrefix, prefixes[1])
	if node.Type != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	}
	fmt.Fprintf(w, "%s%s├── Lbrace= %s\n", parentPrefix, prefixes[1], fs.Position(node.Lbrace))
	fmt.Fprintf(w, "%s%s├── Elts (length=%d)\n", parentPrefix, prefixes[1], len(node.Elts))
	for i := range node.Elts {
		if i < len(node.Elts)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Elts[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Elts[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── Rbrace= %s\n", parentPrefix, prefixes[1], fs.Position(node.Rbrace))
	fmt.Fprintf(w, "%s%s└── Incomplete = %t\n", parentPrefix, prefixes[1], node.Incomplete)
}

func funcLit(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.FuncLit) {
	fmt.Fprintf(w, "%s%sBasicLit\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Type\n", parentPrefix, prefixes[1])
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	fmt.Fprintf(w, "%s%s└── Body\n", parentPrefix, prefixes[1])
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body)
}
