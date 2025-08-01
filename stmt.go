package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func assignStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.AssignStmt) {
	fmt.Fprintf(w, "%s%sAssignStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Lhs (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Lhs))
	for i := range node.Lhs {
		if i < len(node.Lhs)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getMiddlePrefixes(), node.Lhs[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Lhs[i])
		}
	}
	fmt.Fprintf(w, "%s%s%s TokPos = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.TokPos))
	fmt.Fprintf(w, "%s%s%s Tok = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Tok)
	fmt.Fprintf(w, "%s%s%s Rhs (length=%d)\n", parentPrefix, prefixes[1], getTailPrefix(), len(node.Rhs))
	for i := range node.Rhs {
		if i < len(node.Rhs)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, getMiddlePrefixes(), node.Rhs[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Rhs[i])
		}
	}
}

func badStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BadStmt) {
	fmt.Fprintf(w, "%s%sBadStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s From = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.From))
	fmt.Fprintf(w, "%s%s%s To = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.To))
}

func blockStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BlockStmt) {
	fmt.Fprintf(w, "%s%sBlockStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Lbrace= %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lbrace))
	fmt.Fprintf(w, "%s%s%s List (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.List))
	for i := range node.List {
		if i < len(node.List)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getMiddlePrefixes(), node.List[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.List[i])
		}
	}
	fmt.Fprintf(w, "%s%s%s Rbrace= %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rbrace))
}

func branchStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BranchStmt) {
	fmt.Fprintf(w, "%s%sBranchStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s TokPos = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.TokPos))
	fmt.Fprintf(w, "%s%s%s Tok = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Tok)
	fmt.Fprintf(w, "%s%s%s Label\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Label != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Label)
	}
}

func declStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.DeclStmt) {
	fmt.Fprintf(w, "%s%sDeclStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Decl\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Decl)
}

func deferStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.DeferStmt) {
	fmt.Fprintf(w, "%s%sDeferStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Defer = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Defer))
	fmt.Fprintf(w, "%s%s%s Call\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Call)
}

func emptyStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.EmptyStmt) {
	fmt.Fprintf(w, "%s%sEmptyStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Semicolon = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Semicolon))
	fmt.Fprintf(w, "%s%s%s Implicit = %t\n", parentPrefix, prefixes[1], getTailPrefix(), node.Implicit)
}

func exprStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ExprStmt) {
	fmt.Fprintf(w, "%s%sExprStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.X)
}

func forStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ForStmt) {
	fmt.Fprintf(w, "%s%sForStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s For = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.For))
	fmt.Fprintf(w, "%s%s%s Init\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Init != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Init)
	}
	fmt.Fprintf(w, "%s%s%s Cond\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Cond != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Cond)
	}
	fmt.Fprintf(w, "%s%s%s Post\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Post != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Post)
	}
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Body)
}

func goStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.GoStmt) {
	fmt.Fprintf(w, "%s%sGoStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Go = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Go))
	fmt.Fprintf(w, "%s%s%s Call\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Call)
}

func ifStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.IfStmt) {
	fmt.Fprintf(w, "%s%sIfStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s If = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.If))
	fmt.Fprintf(w, "%s%s%s Init\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Init != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Init)
	}
	fmt.Fprintf(w, "%s%s%s Cond\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Cond)
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Body)
	fmt.Fprintf(w, "%s%s%s Else\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Else != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Else)
	}
}

func incDecStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.IncDecStmt) {
	fmt.Fprintf(w, "%s%sIncDecStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s TokPos= %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.TokPos))
	fmt.Fprintf(w, "%s%s%s Tok = %s\n", parentPrefix, prefixes[1], getTailPrefix(), node.Tok)
}

func labeledStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.LabeledStmt) {
	fmt.Fprintf(w, "%s%sLabeledStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Label\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Label)
	fmt.Fprintf(w, "%s%s%s Colon = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Colon))
	fmt.Fprintf(w, "%s%s%s Stmt\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Stmt)
}

func rangeStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.RangeStmt) {
	fmt.Fprintf(w, "%s%sRangeStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s For = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.For))
	fmt.Fprintf(w, "%s%s%s Key\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Key != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Key)
	}
	fmt.Fprintf(w, "%s%s%s Value\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Value != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Value)
	}
	fmt.Fprintf(w, "%s%s%s TokPos = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.TokPos))
	fmt.Fprintf(w, "%s%s%s Tok = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Tok)
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Body)
}

func returnStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ReturnStmt) {
	fmt.Fprintf(w, "%s%sReturnStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Return = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Return))
	fmt.Fprintf(w, "%s%s%s Results\n", parentPrefix, prefixes[1], getTailPrefix())
	for i := range node.Results {
		if i < len(node.Results)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, getMiddlePrefixes(), node.Results[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Results[i])
		}
	}
}

func selectStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.SelectStmt) {
	fmt.Fprintf(w, "%s%sSelectStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Select = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Select))
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Body)
}

func sendStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.SendStmt) {
	fmt.Fprintf(w, "%s%sSendStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Chan\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Chan)
	fmt.Fprintf(w, "%s%s%s Arrow = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Arrow))
	fmt.Fprintf(w, "%s%s%s Value\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Value)
}

func switchStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.SwitchStmt) {
	fmt.Fprintf(w, "%s%sSwitchStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Switch = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Switch))
	fmt.Fprintf(w, "%s%s%s Init\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Init != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Init)
	}
	fmt.Fprintf(w, "%s%s%s Tag\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Tag != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Tag)
	}
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Body)
}

func typeSwitchStmt(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.TypeSwitchStmt) {
	fmt.Fprintf(w, "%s%sTypeSwitchStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Switch = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Switch))
	fmt.Fprintf(w, "%s%s%s Init\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Init != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Init)
	}
	fmt.Fprintf(w, "%s%s%s Assign\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Assign)
	fmt.Fprintf(w, "%s%s%s Body\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Body)
}
