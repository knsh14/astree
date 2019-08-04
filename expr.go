package astree

import (
	"fmt"
	"go/ast"
	"io"
)

func badExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.BadExpr) {
	fmt.Fprintf(w, "%s%sBadExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── From = %v\n", parentPrefix, prefixes[1], node.From)
	fmt.Fprintf(w, "%s%s└── To = %v\n", parentPrefix, prefixes[1], node.To)
}

func binaryExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.BinaryExpr) {
	fmt.Fprintf(w, "%s%sBinaryExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s├── Op = %s\n", parentPrefix, prefixes[1], node.Op)
	fmt.Fprintf(w, "%s%s└── Y\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Y)
}

func callExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.CallExpr) {
	fmt.Fprintf(w, "%s%sCallExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Fun\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Fun)
	fmt.Fprintf(w, "%s%s├── Lparen = %v\n", parentPrefix, prefixes[1], node.Lparen)
	fmt.Fprintf(w, "%s%s├── Args (length=%d)\n", parentPrefix, prefixes[1], len(node.Args))
	for i := range node.Args {
		if i < len(node.Args)-1 {
			Tree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node.Args[i])
		} else {
			Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Args[i])
		}
	}
	fmt.Fprintf(w, "%s%s├── Ellipsis = %v\n", parentPrefix, prefixes[1], node.Ellipsis)
	fmt.Fprintf(w, "%s%s└── Rparen = %v\n", parentPrefix, prefixes[1], node.Rparen)
}

func indexExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.IndexExpr) {
	fmt.Fprintf(w, "%s%sIndexExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s├── Lbrack = %v\n", parentPrefix, prefixes[1], node.Lbrack)
	fmt.Fprintf(w, "%s%s├── Index\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Index)
	fmt.Fprintf(w, "%s%s└── Rbrack = %v\n", parentPrefix, prefixes[1], node.Rbrack)
}

func keyValueExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.KeyValueExpr) {
	fmt.Fprintf(w, "%s%sKeyValueExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Key\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Key)
	fmt.Fprintf(w, "%s%s├── Colon = %v\n", parentPrefix, prefixes[1], node.Colon)
	fmt.Fprintf(w, "%s%s└── Value\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Value)
}

func parenExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.ParenExpr) {
	fmt.Fprintf(w, "%s%sParenExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Lparen = %v\n", parentPrefix, prefixes[1], node.Lparen)
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s└── Rparen = %v\n", parentPrefix, prefixes[1], node.Rparen)
}

func selectorExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.SelectorExpr) {
	fmt.Fprintf(w, "%s%sSelectorExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s└── Sel\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Sel)
}

func sliceExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.SliceExpr) {
	fmt.Fprintf(w, "%s%sSliceExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s├── Lbrack = %v\n", parentPrefix, prefixes[1], node.Lbrack)
	fmt.Fprintf(w, "%s%s├── Low\n", parentPrefix, prefixes[1])
	if node.Low != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Low)
	}
	fmt.Fprintf(w, "%s%s├── High\n", parentPrefix, prefixes[1])
	if node.High != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.High)
	}
	fmt.Fprintf(w, "%s%s├── Max\n", parentPrefix, prefixes[1])
	if node.Max != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Max)
	}
	fmt.Fprintf(w, "%s%s├── Slice3 = %t\n", parentPrefix, prefixes[1], node.Slice3)
	fmt.Fprintf(w, "%s%s└── Rbrack = %v\n", parentPrefix, prefixes[1], node.Rbrack)
}

func starExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.StarExpr) {
	fmt.Fprintf(w, "%s%sStarExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Star = %v\n", parentPrefix, prefixes[1], node.Star)
	fmt.Fprintf(w, "%s%s└── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.X)
}

func typeAssertExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.TypeAssertExpr) {
	fmt.Fprintf(w, "%s%sTypeAssertExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.X)
	fmt.Fprintf(w, "%s%s├── Lparen = %v\n", parentPrefix, prefixes[1], node.Lparen)
	fmt.Fprintf(w, "%s%s├── Type\n", parentPrefix, prefixes[1])
	if node.Type != nil {
		Tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Type)
	} else {
		fmt.Fprintf(w, "%s%sx.(type)\n", parentPrefix+prefixes[1]+middleLine, tailPrefixes[0])
	}
	fmt.Fprintf(w, "%s%s└── Rparen = %v\n", parentPrefix, prefixes[1], node.Rparen)
}

func unaryExpr(w io.Writer, parentPrefix string, prefixes []string, node *ast.UnaryExpr) {
	fmt.Fprintf(w, "%s%sUnaryExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── OpPos = %v\n", parentPrefix, prefixes[1], node.OpPos)
	fmt.Fprintf(w, "%s%s├── Op = %s\n", parentPrefix, prefixes[1], node.Op)
	fmt.Fprintf(w, "%s%s└── X\n", parentPrefix, prefixes[1])
	Tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.X)
}
