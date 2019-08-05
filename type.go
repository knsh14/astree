package astree

import (
	"fmt"
	"go/ast"
	"io"
)

func arrayType(w io.Writer, parentPrefix string, prefixes []string, node *ast.ArrayType) {
	fmt.Fprintf(w, "%s%sArrayType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Lbrack = %v\n", parentPrefix, prefixes[1], node.Lbrack)
	fmt.Fprintf(w, "%s%s├── Len\n", parentPrefix, prefixes[1])
	if node.Len != nil {
		tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Len)
	}
	fmt.Fprintf(w, "%s%s└── Elt\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Elt)
}

func chanType(w io.Writer, parentPrefix string, prefixes []string, node *ast.ChanType) {
	fmt.Fprintf(w, "%s%sChanType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Begin = %v\n", parentPrefix, prefixes[1], node.Begin)
	fmt.Fprintf(w, "%s%s├── Arrow = %v\n", parentPrefix, prefixes[1], node.Arrow)
	fmt.Fprintf(w, "%s%s├── Dir = %v\n", parentPrefix, prefixes[1], node.Dir)
	fmt.Fprintf(w, "%s%s└── Value\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Value)
}

func funcType(w io.Writer, parentPrefix string, prefixes []string, node *ast.FuncType) {
	fmt.Fprintf(w, "%s%sFuncType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Func = %v\n", parentPrefix, prefixes[1], node.Func)
	fmt.Fprintf(w, "%s%s├── Params\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Params)
	fmt.Fprintf(w, "%s%s└── Results\n", parentPrefix, prefixes[1])
	if node.Results != nil {
		tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Results)
	}
}

func interfaceType(w io.Writer, parentPrefix string, prefixes []string, node *ast.InterfaceType) {
	fmt.Fprintf(w, "%s%sInterfaceType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Interface = %v\n", parentPrefix, prefixes[1], node.Interface)
	fmt.Fprintf(w, "%s%s├── Methods\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Methods)
	fmt.Fprintf(w, "%s%s└── Incomplete = %t\n", parentPrefix, prefixes[1], node.Incomplete)
}

func mapType(w io.Writer, parentPrefix string, prefixes []string, node *ast.MapType) {
	fmt.Fprintf(w, "%s%sMapType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Map = %v\n", parentPrefix, prefixes[1], node.Map)
	fmt.Fprintf(w, "%s%s├── Key\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Key)
	fmt.Fprintf(w, "%s%s└── Value\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, node.Value)
}

func structType(w io.Writer, parentPrefix string, prefixes []string, node *ast.StructType) {
	fmt.Fprintf(w, "%s%sStructType\n", parentPrefix, prefixes[0])
	fmt.Fprintf(w, "%s%s├── Struct = %v\n", parentPrefix, prefixes[1], node.Struct)
	fmt.Fprintf(w, "%s%s├── Fields\n", parentPrefix, prefixes[1])
	tree(w, parentPrefix+prefixes[1]+middleLine, tailPrefixes, node.Fields)
	fmt.Fprintf(w, "%s%s└── Incomplete = %t\n", parentPrefix, prefixes[1], node.Incomplete)
}
