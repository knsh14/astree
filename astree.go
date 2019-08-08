package astree

import (
	"go/ast"
	"go/token"
	"io"
)

const (
	middleLine = "│   "
	tailLine   = "    "
)

var (
	middlePrefixes = []string{"├── ", "│   "}
	tailPrefixes   = []string{"└── ", "    "}
)

// File prints AST of one file
func File(w io.Writer, fs *token.FileSet, node *ast.File) {
	tree(w, fs, "", []string{"", ""}, node)
}

// Packages prints AST of result from go/parser.ParseDir
func Packages(w io.Writer, fs *token.FileSet, pkgs map[string]*ast.Package) {
	count := 0
	for k, v := range pkgs {
		if count < len(pkgs)-1 {
			tree(w, fs, "", []string{middlePrefixes[0] + k + ":", middlePrefixes[1]}, v)
		} else {
			tree(w, fs, "", []string{tailPrefixes[0] + k + ":", tailPrefixes[1]}, v)
		}
		count++
	}
}

// Node prints AST node
func Node(w io.Writer, fs *token.FileSet, node ast.Node) {
	tree(w, fs, "", []string{"", ""}, node)
}

// Tree desplays ast nodes like tree
func tree(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node ast.Node) {
	switch n := node.(type) {
	case *ast.File:
		file(w, fs, parentPrefix, prefixes, n)

	// Spec
	case *ast.ImportSpec:
		importSpec(w, fs, parentPrefix, prefixes, n)
	case *ast.ValueSpec:
		valueSpec(w, fs, parentPrefix, prefixes, n)
	case *ast.TypeSpec:
		typeSpec(w, fs, parentPrefix, prefixes, n)

	case *ast.Ident:
		ident(w, fs, parentPrefix, prefixes, n)
	case *ast.BasicLit:
		basicLit(w, fs, parentPrefix, prefixes, n)
	case *ast.CompositeLit:
		compositeLit(w, fs, parentPrefix, prefixes, n)
	case *ast.FuncLit:
		funcLit(w, fs, parentPrefix, prefixes, n)

	// comment
	case *ast.CommentGroup:
		commentGroup(w, fs, parentPrefix, prefixes, n)
	case *ast.Comment:
		comment(w, fs, parentPrefix, prefixes, n)

	// Decls
	case *ast.GenDecl:
		genDecl(w, fs, parentPrefix, prefixes, n)
	case *ast.BadDecl:
		badDecl(w, fs, parentPrefix, prefixes, n)
	case *ast.FuncDecl:
		funcDecl(w, fs, parentPrefix, prefixes, n)

		// Expr
	case *ast.BadExpr:
		badExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.BinaryExpr:
		binaryExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.CallExpr:
		callExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.IndexExpr:
		indexExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.KeyValueExpr:
		keyValueExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.ParenExpr:
		parenExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.SelectorExpr:
		selectorExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.SliceExpr:
		sliceExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.StarExpr:
		starExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.TypeAssertExpr:
		typeAssertExpr(w, fs, parentPrefix, prefixes, n)
	case *ast.UnaryExpr:
		unaryExpr(w, fs, parentPrefix, prefixes, n)

		// Statement
	case *ast.AssignStmt:
		assignStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.BadStmt:
		badStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.BlockStmt:
		blockStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.BranchStmt:
		branchStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.DeclStmt:
		declStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.DeferStmt:
		deferStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.EmptyStmt:
		emptyStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.ExprStmt:
		exprStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.ForStmt:
		forStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.GoStmt:
		goStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.IfStmt:
		ifStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.IncDecStmt:
		incDecStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.LabeledStmt:
		labeledStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.RangeStmt:
		rangeStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.ReturnStmt:
		returnStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.SelectStmt:
		selectStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.SendStmt:
		sendStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.SwitchStmt:
		switchStmt(w, fs, parentPrefix, prefixes, n)
	case *ast.TypeSwitchStmt:
		typeSwitchStmt(w, fs, parentPrefix, prefixes, n)

		// Type
	case *ast.ArrayType:
		arrayType(w, fs, parentPrefix, prefixes, n)
	case *ast.ChanType:
		chanType(w, fs, parentPrefix, prefixes, n)
	case *ast.FuncType:
		funcType(w, fs, parentPrefix, prefixes, n)
	case *ast.InterfaceType:
		interfaceType(w, fs, parentPrefix, prefixes, n)
	case *ast.MapType:
		mapType(w, fs, parentPrefix, prefixes, n)
	case *ast.StructType:
		structType(w, fs, parentPrefix, prefixes, n)

	case *ast.CaseClause:
		caseClause(w, fs, parentPrefix, prefixes, n)
	case *ast.CommClause:
		commClause(w, fs, parentPrefix, prefixes, n)
	case *ast.Ellipsis:
		ellipsis(w, fs, parentPrefix, prefixes, n)
	case *ast.Field:
		field(w, fs, parentPrefix, prefixes, n)
	case *ast.FieldList:
		fieldList(w, fs, parentPrefix, prefixes, n)
	case *ast.Package:
		package2(w, fs, parentPrefix, prefixes, n)
	}
}
