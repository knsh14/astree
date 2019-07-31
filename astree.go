package astree

import (
	"fmt"
	"go/ast"
)

const (
	middleLine = " │    "
	tailLine   = "      "
)

var (
	middlePrefixes = []string{" ├─── ", " │    "}
	tailPrefixes   = []string{" └─── ", "      "}
)

// Tree desplays ast nodes like tree
func Tree(parentPrefix string, prefixes []string, node ast.Node) {
	switch n := node.(type) {
	case *ast.File:
		file(parentPrefix, prefixes, n)
	case *ast.ImportSpec:
		importSpec(parentPrefix, prefixes, n)
	case *ast.Ident:
		ident(parentPrefix, prefixes, n)
	case *ast.BasicLit:
		basicLit(parentPrefix, prefixes, n)
	case *ast.CommentGroup:
		commentGroup(parentPrefix, prefixes, n)
	case *ast.Comment:
		comment(parentPrefix, prefixes, n)
	}
}

func file(parentPrefix string, prefixes []string, node *ast.File) {
	fmt.Printf("%s%sFile\n", parentPrefix, prefixes[0])
	fmt.Printf("%s%s ├── Doc\n", parentPrefix, prefixes[1])
	Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	fmt.Printf("%s%s ├── Name\n", parentPrefix, prefixes[1])
	Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Name)
	fmt.Printf("%s%s ├── Import (length=%d)\n", parentPrefix, prefixes[1], len(node.Imports))
	for i := range node.Imports {
		if i == len(node.Imports)-1 {
			Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Imports[i])
		} else {
			Tree(parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Imports[i])
		}
	}
	fmt.Printf("%s%s ├── Decls (length=%d)\n", parentPrefix, prefixes[1], len(node.Decls))
	for i := range node.Decls {
		if i == len(node.Decls)-1 {
			Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Decls[i])
		} else {
			Tree(parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Decls[i])
		}
	}
}

func importSpec(parentPrefix string, prefixes []string, node *ast.ImportSpec) {
	fmt.Printf("%s%sImportSpec\n", parentPrefix, prefixes[0])
	fmt.Printf("%s%s ├── Doc\n", parentPrefix, prefixes[1])
	Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	fmt.Printf("%s%s ├── Name\n", parentPrefix, prefixes[1])
	Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Name)
	fmt.Printf("%s%s ├── Path\n", parentPrefix, prefixes[1])
	Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Path)
	fmt.Printf("%s%s └── Comment\n", parentPrefix, prefixes[1])
	Tree(parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comment)
}

func ident(parentPrefix string, prefixes []string, node *ast.Ident) {
	if node == nil {
		fmt.Printf("%s%sIdent = nil\n", parentPrefix, prefixes[0])
		return
	}
	fmt.Printf("%s%sIdent\n", parentPrefix, prefixes[0])
	fmt.Printf("%s%s └── Name = %s\n", parentPrefix, prefixes[1], node.Name)
}

func basicLit(parentPrefix string, prefixes []string, node *ast.BasicLit) {
	if node == nil {
		fmt.Printf("%s%sBasicLit = nil\n", parentPrefix, prefixes[0])
		return
	}
	fmt.Printf("%s%sBasicLit\n", parentPrefix, prefixes[0])
	fmt.Printf("%s%s ├── Kind = %s\n", parentPrefix, prefixes[1], node.Kind)
	fmt.Printf("%s%s └── Value = %s\n", parentPrefix, prefixes[1], node.Value)
}

func commentGroup(parentPrefix string, prefixes []string, node *ast.CommentGroup) {
	if node == nil {
		fmt.Printf("%s%sCommentGroup = nil\n", parentPrefix, prefixes[0])
		return
	}
	fmt.Printf("%s%sCommentGroup\n", parentPrefix, prefixes[0])
	for i, comment := range node.List {
		if i == len(node.List)-1 {
			Tree(parentPrefix+prefixes[1]+middleLine, tailPrefixes, comment)
		} else {
			Tree(parentPrefix+prefixes[1]+middleLine, middlePrefixes, comment)
		}
	}
}
func comment(parentPrefix string, prefixes []string, node *ast.Comment) {
	if node == nil {
		return
	}
	fmt.Printf("%s%sComment\n", parentPrefix, prefixes[0])
	fmt.Printf("%s%s └── Text = %s\n", parentPrefix, prefixes[1], node.Text)
}
