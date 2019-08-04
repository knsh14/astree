package astree

import (
	"fmt"
	"go/ast"
	"io"
)

func assignStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.AssignStmt) {
	fmt.Fprintf(w, "%s%sAssignStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Lhs (length=%d)\n", parentPrefix, prefixes[1], len(node.Lhs))
	for i := range node.Lhs {
		if i < len(node.Lhs)-1 {
			Tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Lhs[i])
		} else {
			Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Lhs[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── TokPos = %v\n", parentPrefix, prefixes[1], node.TokPos)
	fmt.Fprintf(w, "%s%s├── Tok = %s\n", parentPrefix, prefixes[1], node.Tok)
	fmt.Fprintf(w, "%s%s└── Rhs (length=%d)\n", parentPrefix, prefixes[1], len(node.Rhs))
	for i := range node.Rhs {
		if i < len(node.Rhs)-1 {
			Tree(w, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Rhs[i])
		} else {
			Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Rhs[i])
		}
	}
}

func badStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.BadStmt) {
	fmt.Fprintf(w, "%s%sBadStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── From = %v\n", parentPrefix, prefixes[1], node.From)
	fmt.Fprintf(w, "%s%s└── To = %v\n", parentPrefix, prefixes[1], node.To)
}

func blockStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.BlockStmt) {
	fmt.Fprintf(w, "%s%sBlockStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Lbrace= %v\n", parentPrefix, prefixes[1], node.Lbrace)
	fmt.Fprintf(w, "%s%s├── List (length=%d)\n", parentPrefix, prefixes[1], len(node.List))
	for i := range node.List {
		if i < len(node.List)-1 {
			Tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.List[i])
		} else {
			Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.List[i])
		}
	}
	fmt.Fprintf(w, "%s%s└── Rbrace= %v\n", parentPrefix, prefixes[1], node.Rbrace)
}

func branchStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.BranchStmt) {
	fmt.Fprintf(w, "%s%sBranchStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── TokPos = %v\n", parentPrefix, prefixes[1], node.TokPos)
	fmt.Fprintf(w, "%s%s├── Tok = %s\n", parentPrefix, prefixes[1], node.Tok)
	fmt.Fprintf(w, "%s%s└── Label\n", parentPrefix, prefixes[1])
	if node.Label != nil {
		Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Label)
	}
}

func declStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.DeclStmt) {
	fmt.Fprintf(w, "%s%sDeclStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s└── Decl\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Decl)
}

func deferStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.DeferStmt) {
	fmt.Fprintf(w, "%s%sDeferStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Defer = %v\n", parentPrefix, prefixes[1], node.Defer)
	fmt.Fprintf(w, "%s%s└── Call\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Call)
}

func emptyStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.EmptyStmt) {
	fmt.Fprintf(w, "%s%sEmptyStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Semicolon = %v\n", parentPrefix, prefixes[1], node.Semicolon)
	fmt.Fprintf(w, "%s%s└── Implicit = %t\n", parentPrefix, prefixes[1], node.Implicit)
}

func exprStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.ExprStmt) {
	fmt.Fprintf(w, "%s%sExprStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s└── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.X)
}

func forStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.ForStmt) {
	fmt.Fprintf(w, "%s%sForStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── For = %v\n", parentPrefix, prefixes[1], node.For)
	fmt.Fprintf(w, "%s%s├── Init\n", parentPrefix, prefixes[1])
	if node.Init != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Init)
	}
	fmt.Fprintf(w, "%s%s├── Cond\n", parentPrefix, prefixes[1])
	if node.Cond != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Cond)
	}
	fmt.Fprintf(w, "%s%s├── Post\n", parentPrefix, prefixes[1])
	if node.Post != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Post)
	}
	fmt.Fprintf(w, "%s%s└── Body\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body)
}

func goStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.GoStmt) {
	fmt.Fprintf(w, "%s%sGoStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Go = %v\n", parentPrefix, prefixes[1], node.Go)
	fmt.Fprintf(w, "%s%s└── Call\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Call)
}

func ifStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.IfStmt) {
	fmt.Fprintf(w, "%s%sIfStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── If = %v\n", parentPrefix, prefixes[1], node.If)
	fmt.Fprintf(w, "%s%s├── Init\n", parentPrefix, prefixes[1])
	if node.Init != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Init)
	}
	fmt.Fprintf(w, "%s%s├── Cond\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Cond)
	fmt.Fprintf(w, "%s%s├── Body\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Body)
	fmt.Fprintf(w, "%s%s└── Else\n", parentPrefix, prefixes[1])
	if node.Else != nil {
		Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Else)
	}
}

func incDecStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.IncDecStmt) {
	fmt.Fprintf(w, "%s%sIncDecStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s├── TokPos= %v\n", parentPrefix, prefixes[1], node.TokPos)
	fmt.Fprintf(w, "%s%s└── Tok = %s\n", parentPrefix, prefixes[1], node.Tok)
}

func labeledStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.LabeledStmt) {
	fmt.Fprintf(w, "%s%sLabeledStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Label\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Label)
	fmt.Fprintf(w, "%s%s├── Colon = %v\n", parentPrefix, prefixes[1], node.Colon)
	fmt.Fprintf(w, "%s%s└── Stmt\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Stmt)
}

func rangeStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.RangeStmt) {
	fmt.Fprintf(w, "%s%sRangeStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── For = %v\n", parentPrefix, prefixes[1], node.For)
	fmt.Fprintf(w, "%s%s├── Key\n", parentPrefix, prefixes[1])
	if node.Key != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Key)
	}
	fmt.Fprintf(w, "%s%s├── Value\n", parentPrefix, prefixes[1])
	if node.Value != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Value)
	}
	fmt.Fprintf(w, "%s%s├── TokPos = %v\n", parentPrefix, prefixes[1], node.TokPos)
	fmt.Fprintf(w, "%s%s├── Tok = %s\n", parentPrefix, prefixes[1], node.Tok)
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s└── Body\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body)
}

func returnStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.ReturnStmt) {
	fmt.Fprintf(w, "%s%sReturnStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Return = %v\n", parentPrefix, prefixes[1], node.Return)
	fmt.Fprintf(w, "%s%s└── Results\n", parentPrefix, prefixes[1])
	for i := range node.Results {
		if i < len(node.Results)-1 {
			Tree(w, parentPrefix+prefixes[1]+tailLine, middlePrefixes, node.Results[i])
		} else {
			Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Results[i])
		}
	}
}

func selectStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.SelectStmt) {
	fmt.Fprintf(w, "%s%sSelectStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Select = %v\n", parentPrefix, prefixes[1], node.Select)
	fmt.Fprintf(w, "%s%s└── Body\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body)
}

func sendStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.SendStmt) {
	fmt.Fprintf(w, "%s%sSendStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Chan\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Chan)
	fmt.Fprintf(w, "%s%s├── Arrow = %v\n", parentPrefix, prefixes[1], node.Arrow)
	fmt.Fprintf(w, "%s%s└── Value\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Value)
}

func switchStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.SwitchStmt) {
	fmt.Fprintf(w, "%s%sSwitchStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Switch = %v\n", parentPrefix, prefixes[1], node.Switch)
	fmt.Fprintf(w, "%s%s├── Init\n", parentPrefix, prefixes[1])
	if node.Init != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Init)
	}
	fmt.Fprintf(w, "%s%s├── Tag\n", parentPrefix, prefixes[1])
	if node.Tag != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Tag)
	}
	fmt.Fprintf(w, "%s%s└── Body\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body)
}

func typeSwitchStmt(w io.Writer, parentPrefix string, prefixes []string, node *ast.TypeSwitchStmt) {
	fmt.Fprintf(w, "%s%sTypeSwitchStmt\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Switch = %v\n", parentPrefix, prefixes[1], node.Switch)
	fmt.Fprintf(w, "%s%s├── Init\n", parentPrefix, prefixes[1])
	if node.Init != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Init)
	}
	fmt.Fprintf(w, "%s%s├── Assign\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Assign)
	fmt.Fprintf(w, "%s%s└── Body\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Body)
}
