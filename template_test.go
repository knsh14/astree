package astree

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIdent(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	err := Initialize(fset)
	t.Run("fixed", func(t *testing.T) {
		expected := `Ident
├── NamePos = -
├── Name = hello
└── Obj
`
		i := &ast.Ident{
			Name: "hello",
		}
		var b bytes.Buffer
		if err != nil {
			t.Fatal(err)
		}
		err = tmplident(&b, "", []string{"", ""}, i)
		if err != nil {
			t.Fatal(err)
		}
		res := b.String()
		if diff := cmp.Diff(expected, res); diff != "" {
			t.Fatalf("(-want +got):\n%s", diff)
		}
	})

	t.Run("back compatible", func(t *testing.T) {
		i := &ast.Ident{
			Name: "hello",
		}
		fset := token.NewFileSet() // positions are relative to fset
		var tb bytes.Buffer
		err := TemplNode(&tb, fset, i)
		if err != nil {
			t.Fatal(err)
		}

		var fb bytes.Buffer
		ident(&fb, fset, "", []string{"", ""}, i)

		if diff := cmp.Diff(tb.String(), fb.String()); diff != "" {
			t.Fatalf("(-want +got):\n%s", diff)
		}
	})
}

func TestObject(t *testing.T) {
	example := `package main
func main() {
	v := 10
	println(v)
}`

	fset := token.NewFileSet() // positions are relative to fset
	err := Initialize(fset)
	v, err := parser.ParseFile(fset, "", example, parser.Mode(0))
	if err != nil {
		t.Fatal(err)
	}

	var actual, expected bytes.Buffer
	c := true
	ast.Inspect(v, func(node ast.Node) bool {
		if n, ok := node.(*ast.Ident); c && ok && n.Name == "v" {
			err = tmplident(&actual, "", []string{"", ""}, n)
			if err != nil {
				t.Fatal(err)
			}
			ident(&expected, fset, "", []string{"", ""}, n)
			c = false
			return c
		}
		return c
	})

	if diff := cmp.Diff(expected.String(), actual.String()); diff != "" {
		t.Fatalf("(-want +got):\n%s", diff)
	}

}
