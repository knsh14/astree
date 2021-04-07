package astree

import (
	"bytes"
	"go/ast"
	"go/token"
	"testing"
)

func TestIdent(t *testing.T) {
	t.Run("fixed", func(t *testing.T) {
		expected := `Ident
├── NamePos = -
├── Name = hello
└── Obj`
		i := &ast.Ident{
			Name: "hello",
		}
		fset := token.NewFileSet() // positions are relative to fset
		var b bytes.Buffer
		err := Initialize(fset)
		if err != nil {
			t.Fatal(err)
		}
		err = tmplident(&b, "", []string{"", ""}, i)
		if err != nil {
			t.Fatal(err)
		}
		res := b.String()
		if res != expected {
			t.Fatalf("result is not expected.\nActual:\n%v\nExpected:\n%s\n", res, expected)
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

		if tb.String() != fb.String() {
			t.Fatalf("result mismatch:\nactual:\n%s\nexpected:\n%s", tb.String(), fb.String())
		}
	})
}
