package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func file(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.File) {
	fmt.Fprintf(w, "%s%sFile\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sDoc\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s%sPackage = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Package))
	fmt.Fprintf(w, "%s%s%sName\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Name)
	fmt.Fprintf(w, "%s%s%sDecls (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Decls))
	for i := range node.Decls {
		if i < len(node.Decls)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Decls[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Decls[i])
		}
	}
	fmt.Fprintf(w, "%s%s%sScope\n", parentPrefix, prefixes[1], getMiddlePrefix())
	fmt.Fprintf(w, "%s%s%sImports (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Imports))
	for i := range node.Imports {
		if i < len(node.Imports)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Imports[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Imports[i])
		}
	}
	fmt.Fprintf(w, "%s%s%sUnresolved (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Unresolved))
	for i := range node.Unresolved {
		if i < len(node.Unresolved)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Unresolved[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Unresolved[i])
		}
	}
	fmt.Fprintf(w, "%s%s%sUnresolved (length=%d)\n", parentPrefix, prefixes[1], getTailPrefix(), len(node.Comments))
	for i := range node.Comments {
		if i < len(node.Comments)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Comments[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comments[i])
		}
	}
}

func ident(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.Ident) {
	fmt.Fprintf(w, "%s%sIdent\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sNamePos = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.NamePos))
	fmt.Fprintf(w, "%s%s%sName = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Name)
	fmt.Fprintf(w, "%s%s%sObj\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Obj != nil {
		object(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Obj)
	}
}

func commentGroup(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.CommentGroup) {
	fmt.Fprintf(w, "%s%sCommentGroup\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sList (length=%d)\n", parentPrefix, prefixes[1], getTailPrefix(), len(node.List))
	for i, comment := range node.List {
		if i == len(node.List)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, comment)
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, comment)
		}
	}
}

func comment(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.Comment) {
	fmt.Fprintf(w, "%s%sComment\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sSlash = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Slash))
	fmt.Fprintf(w, "%s%s%sText = %s\n", parentPrefix, prefixes[1], getTailPrefix(), node.Text)
}

func caseClause(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.CaseClause) {
	fmt.Fprintf(w, "%s%sCaseClause\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sCase = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Case))
	fmt.Fprintf(w, "%s%s%sList (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.List))
	for i := range node.List {
		if i < len(node.List)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.List[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.List[i])
		}
	}
	fmt.Fprintf(w, "%s%s%sColon = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Colon))
	fmt.Fprintf(w, "%s%s%sBody (length=%d)\n", parentPrefix, prefixes[1], getTailPrefix(), len(node.Body))
	for i := range node.Body {
		if i < len(node.Body)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Body[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body[i])
		}
	}
}

func commClause(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.CommClause) {
	fmt.Fprintf(w, "%s%sCommClause\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sCase = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Case))
	fmt.Fprintf(w, "%s%s%sComm\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Comm != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Comm)
	}
	fmt.Fprintf(w, "%s%s%sColon = %v\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Colon)
	fmt.Fprintf(w, "%s%s%sBody (length=%d)\n", parentPrefix, prefixes[1], getTailPrefix(), len(node.Body))
	for i := range node.Body {
		if i == len(node.Body)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Body[i])
		}
	}
}

func ellipsis(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.Ellipsis) {
	fmt.Fprintf(w, "%s%sEllipsis\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sEllipsis = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Ellipsis))
	fmt.Fprintf(w, "%s%s%sElt\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Elt != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Elt)
	}
}

func field(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.Field) {
	fmt.Fprintf(w, "%s%sField\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sDoc\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Doc != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Doc)
	}
	fmt.Fprintf(w, "%s%s%sNames (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Names))
	for i := range node.Names {
		if i < len(node.Names)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Names[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Names[i])
		}
	}
	fmt.Fprintf(w, "%s%s%sType\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	fmt.Fprintf(w, "%s%s%sTag\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Tag != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Tag)
	}
	fmt.Fprintf(w, "%s%s%sComment\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Comment != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Comment)
	}
}

func fieldList(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.FieldList) {
	fmt.Fprintf(w, "%s%sFieldList\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sOpening = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Opening))
	fmt.Fprintf(w, "%s%s%sList (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.List))
	for i := range node.List {
		if i < len(node.List)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.List[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.List[i])
		}
	}
	fmt.Fprintf(w, "%s%s%sClosing = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Closing))
}

func package2(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.Package) {
	fmt.Fprintf(w, "%s%sPackage\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sName = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Name)
	fmt.Fprintf(w, "%s%s%sScope\n", parentPrefix, prefixes[1], getMiddlePrefix())
	fmt.Fprintf(w, "%s%s%sImports (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Imports))
	count := 0
	for k, v := range node.Imports {
		if count < len(node.Imports)-1 {
			object(w, fs, parentPrefix+prefixes[1]+middleLine, []string{middlePrefixes[0] + k + ":", middlePrefixes[1]}, v)
		} else {
			object(w, fs, parentPrefix+prefixes[1]+middleLine, []string{tailPrefixes[0] + k + ":", tailPrefixes[1]}, v)
		}
		count++
	}
	fmt.Fprintf(w, "%s%s%sFiles (length = %d)\n", parentPrefix, prefixes[1], getTailPrefix(), len(node.Files))
	count = 0
	for k, v := range node.Files {
		if count < len(node.Files)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, []string{middlePrefixes[0] + k + ":", middlePrefixes[1]}, v)
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, []string{tailPrefixes[0] + k + ":", tailPrefixes[1]}, v)
		}
		count++
	}
}

func object(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.Object) {
	fmt.Fprintf(w, "%s%sObject\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%sKind = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Kind)
	fmt.Fprintf(w, "%s%s%sName = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Name)
	fmt.Fprintf(w, "%s%s%sDecl = %#v\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Decl)
	fmt.Fprintf(w, "%s%s%sData = %#v\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Data)
	fmt.Fprintf(w, "%s%s%sType = %#v\n", parentPrefix, prefixes[1], getTailPrefix(), node.Type)
}
