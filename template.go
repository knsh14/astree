package astree

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"strings"
	"text/template"
)

var identTemplate2 = `Ident
├── NamePos = {{ position .NamePos }}
├── Name = {{ .Name }}
{{ with .Obj }}└── Obj = {{ obj . }}{{ else }}└── Obj{{ end }}
`
var (
	identTemplate *template.Template
)

func TemplNode(w io.Writer, fs *token.FileSet, node ast.Node) error {
	err := Initialize(fs)
	if err != nil {
		return fmt.Errorf("initialize: %w", err)
	}
	err = tmpltree(w, "", []string{"", ""}, node)
	if err != nil {
		return fmt.Errorf("execute tree: %w", err)
	}
	return nil
}

func Initialize(fs *token.FileSet) error {
	v := strings.Split(identTemplate2, "\n")
	v = append([]string{""}, v...)
	s := "{{ template \"prefix1\" }}" + v[0] + strings.Join(v[1:len(v)-1], "\n"+"{{ template \"prefix1\" }}") + "\n{{ template \"prefix2\" }}" + v[len(v)-1]
	funcMap := template.FuncMap{
		"tree": func(node ast.Node) string {
			return ""
		},
		"position": func(p token.Pos) token.Position {
			return fs.Position(p)
		},
		"obj": func(any interface{}) string {
			return ""
		},
	}

	tmpl, err := template.New("ident").Funcs(funcMap).Parse(s)
	if err != nil {
		return fmt.Errorf("parse ident template: %w", err)
	}

	identTemplate = tmpl
	return nil
}

func tmplident(w io.Writer, parentPrefix string, prefixes []string, node *ast.Ident) error {
	t, err := identTemplate.Clone()
	if err != nil {
		return fmt.Errorf("clone base ident template: %w", err)
	}
	funcMap := template.FuncMap{
		"tree": func(node ast.Node) string {
			tmpltree(w, parentPrefix+prefixes[0], middlePrefixes, node)
			return ""
		},
		"obj": func(obj *ast.Object) string {
			return ""
		},
	}
	t, err = t.Funcs(funcMap).Parse(fmt.Sprintf("{{define \"prefix1\"}}%s{{end}}{{ define \"prefix2\" }}%s{{end}}", parentPrefix+prefixes[0], parentPrefix+prefixes[1]))
	if err != nil {
		return fmt.Errorf("parse prefix templates: %w", err)
	}
	err = t.Execute(w, node)
	if err != nil {
		return fmt.Errorf("execute ident template: %w", err)
	}
	return nil
}

func tmplobject(w io.Writer, parentPrefix string, prefixes []string, node *ast.Object) {
}

func tmpltree(w io.Writer, parentPrefix string, prefixes []string, node ast.Node) error {
	switch n := node.(type) {
	case *ast.Ident:
		err := tmplident(w, parentPrefix, prefixes, n)
		if err != nil {
			return fmt.Errorf("execute tree ident template: %w", err)
		}
	}
	return nil
}

// テンプレート文字列に対して先頭に罫線用のプレフィックスをつける
// template を全部ブロックにする
// 関数を定義する
//
