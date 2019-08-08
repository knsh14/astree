package astree

import (
	"bytes"
	"go/parser"
	"go/token"
	"testing"
)

func TestFile(t *testing.T) {
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
