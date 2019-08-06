package astree

import (
	"go/ast"
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
func File(w io.Writer, node *ast.File) {
	tree(w, "", []string{"", ""}, node)
}

// Packages prints AST of result from go/parser.ParseDir
func Packages(w io.Writer, pkgs map[string]*ast.Package) {
	count := 0
	for k, v := range pkgs {
		if count < len(pkgs)-1 {
			tree(w, "", []string{middlePrefixes[0] + k + ":", middlePrefixes[1]}, v)
		} else {
			tree(w, "", []string{tailPrefixes[0] + k + ":", tailPrefixes[1]}, v)
		}
		count++
	}
}

// Node prints AST node
func Node(w io.Writer, node *ast.File) {
	tree(w, "", []string{"", ""}, node)
}

// Tree desplays ast nodes like tree
func tree(w io.Writer, parentPrefix string, prefixes []string, node ast.Node) {
	switch n := node.(type) {
	case *ast.File:
		file(w, parentPrefix, prefixes, n)

	// Spec
	case *ast.ImportSpec:
		importSpec(w, parentPrefix, prefixes, n)
	case *ast.ValueSpec:
		valueSpec(w, parentPrefix, prefixes, n)
	case *ast.TypeSpec:
		typeSpec(w, parentPrefix, prefixes, n)

	case *ast.Ident:
		ident(w, parentPrefix, prefixes, n)
	case *ast.BasicLit:
		basicLit(w, parentPrefix, prefixes, n)
	case *ast.CompositeLit:
		compositeLit(w, parentPrefix, prefixes, n)
	case *ast.FuncLit:
		funcLit(w, parentPrefix, prefixes, n)

	// comment
	case *ast.CommentGroup:
		commentGroup(w, parentPrefix, prefixes, n)
	case *ast.Comment:
		comment(w, parentPrefix, prefixes, n)

	// Decls
	case *ast.GenDecl:
		genDecl(w, parentPrefix, prefixes, n)
	case *ast.BadDecl:
		badDecl(w, parentPrefix, prefixes, n)
	case *ast.FuncDecl:
		funcDecl(w, parentPrefix, prefixes, n)

		// Expr
	case *ast.BadExpr:
		badExpr(w, parentPrefix, prefixes, n)
	case *ast.BinaryExpr:
		binaryExpr(w, parentPrefix, prefixes, n)
	case *ast.CallExpr:
		callExpr(w, parentPrefix, prefixes, n)
	case *ast.IndexExpr:
		indexExpr(w, parentPrefix, prefixes, n)
	case *ast.KeyValueExpr:
		keyValueExpr(w, parentPrefix, prefixes, n)
	case *ast.ParenExpr:
		parenExpr(w, parentPrefix, prefixes, n)
	case *ast.SelectorExpr:
		selectorExpr(w, parentPrefix, prefixes, n)
	case *ast.SliceExpr:
		sliceExpr(w, parentPrefix, prefixes, n)
	case *ast.StarExpr:
		starExpr(w, parentPrefix, prefixes, n)
	case *ast.TypeAssertExpr:
		typeAssertExpr(w, parentPrefix, prefixes, n)
	case *ast.UnaryExpr:
		unaryExpr(w, parentPrefix, prefixes, n)

		// Statement
	case *ast.AssignStmt:
		assignStmt(w, parentPrefix, prefixes, n)
	case *ast.BadStmt:
		badStmt(w, parentPrefix, prefixes, n)
	case *ast.BlockStmt:
		blockStmt(w, parentPrefix, prefixes, n)
	case *ast.BranchStmt:
		branchStmt(w, parentPrefix, prefixes, n)
	case *ast.DeclStmt:
		declStmt(w, parentPrefix, prefixes, n)
	case *ast.DeferStmt:
		deferStmt(w, parentPrefix, prefixes, n)
	case *ast.EmptyStmt:
		emptyStmt(w, parentPrefix, prefixes, n)
	case *ast.ExprStmt:
		exprStmt(w, parentPrefix, prefixes, n)
	case *ast.ForStmt:
		forStmt(w, parentPrefix, prefixes, n)
	case *ast.GoStmt:
		goStmt(w, parentPrefix, prefixes, n)
	case *ast.IfStmt:
		ifStmt(w, parentPrefix, prefixes, n)
	case *ast.IncDecStmt:
		incDecStmt(w, parentPrefix, prefixes, n)
	case *ast.LabeledStmt:
		labeledStmt(w, parentPrefix, prefixes, n)
	case *ast.RangeStmt:
		rangeStmt(w, parentPrefix, prefixes, n)
	case *ast.ReturnStmt:
		returnStmt(w, parentPrefix, prefixes, n)
	case *ast.SelectStmt:
		selectStmt(w, parentPrefix, prefixes, n)
	case *ast.SendStmt:
		sendStmt(w, parentPrefix, prefixes, n)
	case *ast.SwitchStmt:
		switchStmt(w, parentPrefix, prefixes, n)
	case *ast.TypeSwitchStmt:
		typeSwitchStmt(w, parentPrefix, prefixes, n)

		// Type
	case *ast.ArrayType:
		arrayType(w, parentPrefix, prefixes, n)
	case *ast.ChanType:
		chanType(w, parentPrefix, prefixes, n)
	case *ast.FuncType:
		funcType(w, parentPrefix, prefixes, n)
	case *ast.InterfaceType:
		interfaceType(w, parentPrefix, prefixes, n)
	case *ast.MapType:
		mapType(w, parentPrefix, prefixes, n)
	case *ast.StructType:
		structType(w, parentPrefix, prefixes, n)

	case *ast.CaseClause:
		caseClause(w, parentPrefix, prefixes, n)
	case *ast.CommClause:
		commClause(w, parentPrefix, prefixes, n)
	case *ast.Ellipsis:
		ellipsis(w, parentPrefix, prefixes, n)
	case *ast.Field:
		field(w, parentPrefix, prefixes, n)
	case *ast.FieldList:
		fieldList(w, parentPrefix, prefixes, n)
	case *ast.Package:
		package2(w, parentPrefix, prefixes, n)
	}
}
