package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func genDecl(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.GenDecl) {
	fmt.Fprintf(w, "%s%sGenDecl\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Doc\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Doc)
	}
	fmt.Fprintf(w, "%s%s%s Tok = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Tok)
	fmt.Fprintf(w, "%s%s%s Specs (length=%d)\n", parentPrefix, prefixes[1], getTailPrefix(), len(node.Specs))
	for i := range node.Specs {
		if i < len(node.Specs)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, getMiddlePrefixes(), node.Specs[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Specs[i])
		}
	}
}

func badDecl(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BadDecl) {
	fmt.Fprintf(w, "%s%sBadDecl\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s From = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.From))
	fmt.Fprintf(w, "%s%s%s To = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.To))
}

func funcDecl(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.FuncDecl) {
	fmt.Fprintf(w, "%s%sFuncDecl\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Doc\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Doc)
	}
	fmt.Fprintf(w, "%s%s%s Recv\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Recv != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Recv)
	}
	fmt.Fprintf(w, "%s%s%s Name\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Name)
	fmt.Fprintf(w, "%s%s%s Type\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Type)
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Body != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Body)
	}
}
