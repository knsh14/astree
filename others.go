package astree

import (
	"fmt"
	"go/ast"
	"io"
)

func file(w io.Writer, parentPrefix string, prefixes []string, node *ast.File) {
	fmt.Fprintf(w, "%s%sFile\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Doc\n", parentPrefix, prefixes[1])
	if node.Doc != nil {
		tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s├── Package = %v\n", parentPrefix, prefixes[1], node.Package)
	fmt.Fprintf(w, "%s%s├── Name\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Name)
	fmt.Fprintf(w, "%s%s├── Decls (length=%d)\n", parentPrefix, prefixes[1], len(node.Decls))
	for i := range node.Decls {
		if i < len(node.Decls)-1 {
			tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Decls[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Decls[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── Scope\n", parentPrefix, prefixes[1])
	fmt.Fprintf(w, "%s%s├── Imports (length=%d)\n", parentPrefix, prefixes[1], len(node.Imports))
	for i := range node.Imports {
		if i < len(node.Imports)-1 {
			tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Imports[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Imports[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── Unresolved (length=%d)\n", parentPrefix, prefixes[1], len(node.Unresolved))
	for i := range node.Unresolved {
		if i < len(node.Unresolved)-1 {
			tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Unresolved[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Unresolved[i])
		}
	}
	fmt.Fprintf(w, "%s%s└── Unresolved (length=%d)\n", parentPrefix, prefixes[1], len(node.Comments))
	for i := range node.Comments {
		if i < len(node.Comments)-1 {
			tree(w, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Comments[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comments[i])
		}
	}
}

func ident(w io.Writer, parentPrefix string, prefixes []string, node *ast.Ident) {
	fmt.Fprintf(w, "%s%sIdent\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── NamePos = %v\n", parentPrefix, prefixes[1], node.NamePos)
	fmt.Fprintf(w, "%s%s├── Name = %s\n", parentPrefix, prefixes[1], node.Name)
	fmt.Fprintf(w, "%s%s└── Obj\n", parentPrefix, prefixes[1])
	if node.Obj != nil {
		object(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Obj)
	}
	// tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Obj)
}

func commentGroup(w io.Writer, parentPrefix string, prefixes []string, node *ast.CommentGroup) {
	fmt.Fprintf(w, "%s%sCommentGroup\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── List (length=%d)\n", parentPrefix, prefixes[1], len(node.List))
	for i, comment := range node.List {
		if i == len(node.List)-1 {
			tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, comment)
		} else {
			tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, comment)
		}
	}
}

func comment(w io.Writer, parentPrefix string, prefixes []string, node *ast.Comment) {
	fmt.Fprintf(w, "%s%sComment\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Slash = %v\n", parentPrefix, prefixes[1], node.Slash)
	fmt.Fprintf(w, "%s%s└── Text = %s\n", parentPrefix, prefixes[1], node.Text)
}

func caseClause(w io.Writer, parentPrefix string, prefixes []string, node *ast.CaseClause) {
	fmt.Fprintf(w, "%s%sCaseClause\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Case = %v\n", parentPrefix, prefixes[1], node.Case)
	fmt.Fprintf(w, "%s%s├── List (length=%d)\n", parentPrefix, prefixes[1], len(node.List))
	for i := range node.List {
		if i < len(node.List)-1 {
			tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.List[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.List[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── Colon = %v\n", parentPrefix, prefixes[1], node.Colon)
	fmt.Fprintf(w, "%s%s└── Body (length=%d)\n", parentPrefix, prefixes[1], len(node.Body))
	for i := range node.Body {
		if i < len(node.Body)-1 {
			tree(w, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Body[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body[i])
		}
	}
}

func commClause(w io.Writer, parentPrefix string, prefixes []string, node *ast.CommClause) {
	fmt.Fprintf(w, "%s%sCommClause\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Case = %v\n", parentPrefix, prefixes[1], node.Case)
	fmt.Fprintf(w, "%s%s├── Comm\n", parentPrefix, prefixes[1])
	if node.Comm != nil {
		tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Comm)
	}
	fmt.Fprintf(w, "%s%s├── Colon = %v\n", parentPrefix, prefixes[1], node.Colon)
	fmt.Fprintf(w, "%s%s└── Body (length=%d)\n", parentPrefix, prefixes[1], len(node.Body))
	for i := range node.Body {
		if i == len(node.Body)-1 {
			tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Body[i])
		}
	}
}

func ellipsis(w io.Writer, parentPrefix string, prefixes []string, node *ast.Ellipsis) {
	fmt.Fprintf(w, "%s%sEllipsis\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Ellipsis = %v\n", parentPrefix, prefixes[1], node.Ellipsis)
	fmt.Fprintf(w, "%s%s└── Elt\n", parentPrefix, prefixes[1])
	if node.Elt != nil {
		tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Elt)
	}
}

func field(w io.Writer, parentPrefix string, prefixes []string, node *ast.Field) {
	fmt.Fprintf(w, "%s%sField\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Doc\n", parentPrefix, prefixes[1])
	if node.Doc != nil {
		tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s├── Names (length=%d)\n", parentPrefix, prefixes[1], len(node.Names))
	for i := range node.Names {
		if i < len(node.Names)-1 {
			tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Names[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Names[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── Type\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	fmt.Fprintf(w, "%s%s├── Tag\n", parentPrefix, prefixes[1])
	if node.Tag != nil {
		tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Tag)
	}
	fmt.Fprintf(w, "%s%s└── Comment\n", parentPrefix, prefixes[1])
	if node.Comment != nil {
		tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comment)
	}
}

func fieldList(w io.Writer, parentPrefix string, prefixes []string, node *ast.FieldList) {
	fmt.Fprintf(w, "%s%sFieldList\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Opening = %v\n", parentPrefix, prefixes[1], node.Opening)
	fmt.Fprintf(w, "%s%s├── List (length=%d)\n", parentPrefix, prefixes[1], len(node.List))
	for i := range node.List {
		if i < len(node.List)-1 {
			tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.List[i])
		} else {
			tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.List[i])
		}
	}
	fmt.Fprintf(w, "%s%s└── Closing = %v\n", parentPrefix, prefixes[1], node.Closing)
}

func package2(w io.Writer, parentPrefix string, prefixes []string, node *ast.Package) {
	fmt.Fprintf(w, "%s%sPackage\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Name = %s\n", parentPrefix, prefixes[1], node.Name)
	fmt.Fprintf(w, "%s%s├── Scope\n", parentPrefix, prefixes[1])
	fmt.Fprintf(w, "%s%s├── Imports (length=%d)\n", parentPrefix, prefixes[1], len(node.Imports))
	count := 0
	for k, v := range node.Imports {
		if count < len(node.Imports)-1 {
			object(w, parentPrefix+prefixes[1]+middleLine, []string{middlePrefixes[0] + k + ":", middlePrefixes[1]}, v)
		} else {
			object(w, parentPrefix+prefixes[1]+middleLine, []string{tailPrefixes[0] + k + ":", tailPrefixes[1]}, v)
		}
		count++
	}
	fmt.Fprintf(w, "%s%s└── Files (length = %d)\n", parentPrefix, prefixes[1], len(node.Files))
	count = 0
	for k, v := range node.Files {
		if count < len(node.Files)-1 {
			tree(w, parentPrefix+prefixes[1]+tailLine, []string{middlePrefixes[0] + k + ":", middlePrefixes[1]}, v)
		} else {
			tree(w, parentPrefix+prefixes[1]+tailLine, []string{tailPrefixes[0] + k + ":", tailPrefixes[1]}, v)
		}
		count++
	}
}

func object(w io.Writer, parentPrefix string, prefixes []string, node *ast.Object) {
	fmt.Fprintf(w, "%s%sObject\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Kind = %s\n", parentPrefix, prefixes[1], node.Kind)
	fmt.Fprintf(w, "%s%s├── Name = %s\n", parentPrefix, prefixes[1], node.Name)
	fmt.Fprintf(w, "%s%s├── Decl = %#v\n", parentPrefix, prefixes[1], node.Decl)
	fmt.Fprintf(w, "%s%s├── Data = %#v\n", parentPrefix, prefixes[1], node.Data)
	fmt.Fprintf(w, "%s%s└── Type = %#v\n", parentPrefix, prefixes[1], node.Type)
}
