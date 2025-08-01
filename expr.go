package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func badExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BadExpr) {
	fmt.Fprintf(w, "%s%sBadExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s From = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.From))
	fmt.Fprintf(w, "%s%s%s To = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.To))
}

func binaryExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.BinaryExpr) {
	fmt.Fprintf(w, "%s%sBinaryExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Op = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Op)
	fmt.Fprintf(w, "%s%s%s Y\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Y)
}

func callExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.CallExpr) {
	fmt.Fprintf(w, "%s%sCallExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Fun\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Fun)
	fmt.Fprintf(w, "%s%s%s Lparen = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lparen))
	fmt.Fprintf(w, "%s%s%s Args (length=%d)\n", parentPrefix, prefixes[1], getMiddlePrefix(), len(node.Args))
	for i := range node.Args {
		if i < len(node.Args)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getMiddlePrefixes(), node.Args[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Args[i])
		}
	}
	fmt.Fprintf(w, "%s%s%s Ellipsis = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Ellipsis))
	fmt.Fprintf(w, "%s%s%s Rparen = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rparen))
}

func indexExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.IndexExpr) {
	fmt.Fprintf(w, "%s%sIndexExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Lbrack = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lbrack))
	fmt.Fprintf(w, "%s%s%s Index\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Index)
	fmt.Fprintf(w, "%s%s%s Rbrack = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rbrack))
}

func keyValueExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.KeyValueExpr) {
	fmt.Fprintf(w, "%s%sKeyValueExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Key\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Key)
	fmt.Fprintf(w, "%s%s%s Colon = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Colon))
	fmt.Fprintf(w, "%s%s%s Value\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Value)
}

func parenExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ParenExpr) {
	fmt.Fprintf(w, "%s%sParenExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Lparen = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lparen))
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Rparen = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rparen))
}

func selectorExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.SelectorExpr) {
	fmt.Fprintf(w, "%s%sSelectorExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Sel\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Sel)
}

func sliceExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.SliceExpr) {
	fmt.Fprintf(w, "%s%sSliceExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Lbrack = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lbrack))
	fmt.Fprintf(w, "%s%s%s Low\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Low != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Low)
	}
	fmt.Fprintf(w, "%s%s%s High\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.High != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.High)
	}
	fmt.Fprintf(w, "%s%s%s Max\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Max != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Max)
	}
	fmt.Fprintf(w, "%s%s%s Slice3 = %t\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Slice3)
	fmt.Fprintf(w, "%s%s%s Rbrack = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rbrack))
}

func starExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.StarExpr) {
	fmt.Fprintf(w, "%s%sStarExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Star = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Star))
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.X)
}

func typeAssertExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.TypeAssertExpr) {
	fmt.Fprintf(w, "%s%sTypeAssertExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Lparen = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lparen))
	fmt.Fprintf(w, "%s%s%s Type\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Type != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Type)
	} else {
		fmt.Fprintf(w, "%s%sx.(type)\n", parentPrefix+prefixes[1]+middleLine, getTailPrefixes()[0])
	}
	fmt.Fprintf(w, "%s%s%s Rparen = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rparen))
}

func unaryExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.UnaryExpr) {
	fmt.Fprintf(w, "%s%sUnaryExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s OpPos = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.OpPos))
	fmt.Fprintf(w, "%s%s%s Op = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Op)
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.X)
}

func indexListExpr(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.IndexListExpr) {
	fmt.Fprintf(w, "%s%sIndexListExpr\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s X\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.X)
	fmt.Fprintf(w, "%s%s%s Lbrack = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lbrack))
	fmt.Fprintf(w, "%s%s%s Indices\n", parentPrefix, prefixes[1], getMiddlePrefix())
	for i := range node.Indices {
		if i < len(node.Indices)-1 {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getMiddlePrefixes(), node.Indices[i])
		} else {
			tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Indices[i])
		}
	}
	fmt.Fprintf(w, "%s%s%s Rbrack = %s\n", parentPrefix, prefixes[1], getTailPrefix(), fs.Position(node.Rbrack))
}
