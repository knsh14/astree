package astree

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestFile(t *testing.T) {
	src := `
package main

import "fmt"

var foo int

func main() {
	f.Println("hello world")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}
	Tree("", []string{"", ""}, f)
}
