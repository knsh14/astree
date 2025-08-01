package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func basicLit(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BasicLit) {
	fmt.Fprintf(w, "%s%sBasicLit\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Kind = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Kind)
	fmt.Fprintf(w, "%s%s%s Value = %s\n", parentPrefix, prefixes[1], getTailPrefix(), node.Value)
}

func compositeLit(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.CompositeLit) {
	fmt.Fprintf(w, "%s%sCompositeLit\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Type\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Type != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Type)
	}
	fmt.Fprintf(w, "%s%s%s Lbrace= %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lbrace))
	fmt.Fprintf(w, "%s%s%s Elts (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Elts))
	for i := range node.Elts {
		if i < len(node.Elts)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getMiddlePrefixes(), node.Elts[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Elts[i])
		}
	}
	fmt.Fprintf(w, "%s%s%s Rbrace= %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rbrace))
}

func funcLit(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.FuncLit) {
	fmt.Fprintf(w, "%s%sFuncLit\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Type\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Type)
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Body)
}
