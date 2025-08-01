package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func importSpec(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ImportSpec) {
	fmt.Fprintf(w, "%s%sImportSpec\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Doc\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Doc)
	}
	fmt.Fprintf(w, "%s%s%s Name\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Name != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Name)
	}
	fmt.Fprintf(w, "%s%s%s Path\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Path)
	fmt.Fprintf(w, "%s%s%s Comment\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Comment != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Comment)
	}
	fmt.Fprintf(w, "%s%s%s EndPos = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.EndPos))
}

func typeSpec(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.TypeSpec) {
	fmt.Fprintf(w, "%s%sTypeSpec\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Doc\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Doc)
	}
	fmt.Fprintf(w, "%s%s%s Name\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Name)
	fmt.Fprintf(w, "%s%s%s TypeParams\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.TypeParams != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.TypeParams)
	}
	fmt.Fprintf(w, "%s%s%s Assign = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Assign))
	fmt.Fprintf(w, "%s%s%s Type\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Type)
	fmt.Fprintf(w, "%s%s%s Comment\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Comment != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Comment)
	}
}

func valueSpec(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ValueSpec) {
	fmt.Fprintf(w, "%s%sValueSpec\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Doc\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Doc)
	}
	fmt.Fprintf(w, "%s%s%s Names (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Names))
	for i := range node.Names {
		if i == len(node.Names)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Names[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getMiddlePrefixes(), node.Names[i])
		}
	}
	fmt.Fprintf(w, "%s%s%s Type\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Type != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Type)
	}
	fmt.Fprintf(w, "%s%s%s Values (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Values))
	for i := range node.Values {
		if i < len(node.Values)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getMiddlePrefixes(), node.Values[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Values[i])
		}
	}
	fmt.Fprintf(w, "%s%s%s Comment\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Comment != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Comment)
	}
}
