package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func arrayType(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ArrayType) {
	fmt.Fprintf(w, "%s%sArrayType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Lbrack = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Lbrack))
	fmt.Fprintf(w, "%s%s%s Len\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.Len != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Len)
	}
	fmt.Fprintf(w, "%s%s%s Elt\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Elt)
}

func chanType(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.ChanType) {
	fmt.Fprintf(w, "%s%sChanType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Begin = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Begin))
	fmt.Fprintf(w, "%s%s%s Arrow = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Arrow))
	fmt.Fprintf(w, "%s%s%s Dir = %v\n", parentPrefix, prefixes[1], getMiddlePrefix(), node.Dir)
	fmt.Fprintf(w, "%s%s%s Value\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Value)
}

func funcType(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.FuncType) {
	fmt.Fprintf(w, "%s%sFuncType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Func = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Func))
	fmt.Fprintf(w, "%s%s%s TypeParams\n", parentPrefix, prefixes[1], getMiddlePrefix())
	if node.TypeParams != nil {
		tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.TypeParams)
	}
	fmt.Fprintf(w, "%s%s%s Params\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Params)
	fmt.Fprintf(w, "%s%s%s Results\n", parentPrefix, prefixes[1], getTailPrefix())
	if node.Results != nil {
		tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Results)
	}
}

func interfaceType(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.InterfaceType) {
	fmt.Fprintf(w, "%s%sInterfaceType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Interface = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Interface))
	fmt.Fprintf(w, "%s%s%s Methods\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Methods)
	fmt.Fprintf(w, "%s%s%s Incomplete = %t\n", parentPrefix, prefixes[1], getTailPrefix(), node.Incomplete)
}

func mapType(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.MapType) {
	fmt.Fprintf(w, "%s%sMapType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Map = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Map))
	fmt.Fprintf(w, "%s%s%s Key\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Key)
	fmt.Fprintf(w, "%s%s%s Value\n", parentPrefix, prefixes[1], getTailPrefix())
	tree(w, fs, parentPrefix+prefixes[1]+tailLine, getTailPrefixes(), node.Value)
}

func structType(w io.Writer, fs *token.FileSet, parentPrefix string, prefixes []string, node *ast.StructType) {
	fmt.Fprintf(w, "%s%sStructType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s%s Struct = %s\n", parentPrefix, prefixes[1], getMiddlePrefix(), fs.Position(node.Struct))
	fmt.Fprintf(w, "%s%s%s Fields\n", parentPrefix, prefixes[1], getMiddlePrefix())
	tree(w, fs, parentPrefix+prefixes[1]+middleLine, getTailPrefixes(), node.Fields)
	fmt.Fprintf(w, "%s%s%s Incomplete = %t\n", parentPrefix, prefixes[1], getTailPrefix(), node.Incomplete)
}
