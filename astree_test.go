package astree

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestInternalTree(t *testing.T) {
	src := `package main

import "fmt"

var foo int

func main() {
	foo = 1
	fmt.Println(foo)
}
`

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	b := &bytes.Buffer{}
	tree(b, fset, "", []string{"", ""}, f)
	// its ok if no panic by invalid memory address error
}

func TestFile_NilFileSet(t *testing.T) {
	src := `package main

import "fmt"

var foo int

func main() {
	foo = 1
	fmt.Println(foo)
}
`

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	b := &bytes.Buffer{}
	err = File(b, nil, f)
	if err == nil {
		t.Error("expected to cause error")
	}
	if !strings.Contains(err.Error(), "*token.FileSet is nil") {
		t.Errorf("error message is not expected. actual=%s", err.Error())
	}
}

func TestIdent(t *testing.T) {
	i := &ast.Ident{
		Name: "hello",
	}
	fset := token.NewFileSet() // positions are relative to fset
	b := &bytes.Buffer{}
	ident(b, fset, "", []string{"", ""}, i)
	expect := `Ident
├── NamePos = -
├── Name = hello
└── Obj = nil
`
	res := b.String()
	if res != expect {
		t.Fatalf("result is not expected.\nActual:\n%s\nExpected:\n%s\n", res, expect)
	}
}
