package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func genDecl(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.GenDecl) {
	fmt.Fprintf(w, "%s%sGenDecl\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Doc\n", parentPrefix, prefixes[1])
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s├── Tok = %s\n", parentPrefix, prefixes[1], node.Tok)
	fmt.Fprintf(w, "%s%s└── Specs (length=%d)\n", parentPrefix, prefixes[1], len(node.Specs))
	for i := range node.Specs {
		if i < len(node.Specs)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Specs[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Specs[i])
		}
	}
}

func badDecl(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BadDecl) {
	fmt.Fprintf(w, "%s%sBadDecl\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── From = %s\n", parentPrefix, prefixes[1], fs.Position(node.From))
	fmt.Fprintf(w, "%s%s└── To = %s\n", parentPrefix, prefixes[1], fs.Position(node.To))
}

func funcDecl(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.FuncDecl) {
	fmt.Fprintf(w, "%s%sFuncDecl\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Doc\n", parentPrefix, prefixes[1])
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s├── Recv\n", parentPrefix, prefixes[1])
	if node.Recv != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Recv)
	}
	fmt.Fprintf(w, "%s%s├── Name\n", parentPrefix, prefixes[1])
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Name)
	fmt.Fprintf(w, "%s%s├── Type\n", parentPrefix, prefixes[1])
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	fmt.Fprintf(w, "%s%s└── Body\n", parentPrefix, prefixes[1])
	if node.Body != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body)
	}
}
