package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func importSpec(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ImportSpec) {
	fmt.Fprintf(w, "%s%sImportSpec\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Doc\n", parentPrefix, prefixes[1])
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s├── Name\n", parentPrefix, prefixes[1])
	if node.Name != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Name)
	}
	fmt.Fprintf(w, "%s%s├── Path\n", parentPrefix, prefixes[1])
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Path)
	fmt.Fprintf(w, "%s%s├── Comment\n", parentPrefix, prefixes[1])
	if node.Comment != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comment)
	}
	fmt.Fprintf(w, "%s%s└── EndPos = %s\n", parentPrefix, prefixes[1], fs.Position(node.EndPos))
}

func typeSpec(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.TypeSpec) {
	fmt.Fprintf(w, "%s%sTypeSpec\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Doc\n", parentPrefix, prefixes[1])
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s├── Name\n", parentPrefix, prefixes[1])
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Name)
	fmt.Fprintf(w, "%s%s├── Assign = %s\n", parentPrefix, prefixes[1], fs.Position(node.Assign))
	fmt.Fprintf(w, "%s%s├── Type\n", parentPrefix, prefixes[1])
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	fmt.Fprintf(w, "%s%s└── Comment\n", parentPrefix, prefixes[1])
	if node.Comment != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comment)
	}
}

func valueSpec(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ValueSpec) {
	fmt.Fprintf(w, "%s%sValueSpec\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Doc\n", parentPrefix, prefixes[1])
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s├── Names (length=%d)\n", parentPrefix, prefixes[1], len(node.Names))
	for i := range node.Names {
		if i == len(node.Names)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Names[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Names[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── Type\n", parentPrefix, prefixes[1])
	if node.Type != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	}
	fmt.Fprintf(w, "%s%s├── Values (length=%d)\n", parentPrefix, prefixes[1], len(node.Values))
	for i := range node.Values {
		if i < len(node.Values)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Values[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Values[i])
		}
	}
	fmt.Fprintf(w, "%s%s└── Comment\n", parentPrefix, prefixes[1])
	if node.Comment != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comment)
	}
}
