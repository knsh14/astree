package astree

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"strings"
	"text/template"

	"github.com/morikuni/failure"
)

type templatePair struct {
	Tpl      string
	Template *template.Template
}

var (
	tmplMap = map[string]*templatePair{
		"*ast.Ident": {
			Tpl: `Ident
├── NamePos = {{ position .NamePos }}
├── Name = {{ .Name }}
└── Obj
{{ with .Obj }}{{ obj . }}{{ end }}`,
		},
		"*ast.Object": {
			Tpl: `Object
├── Kind = {{ .Kind }}
├── Name = {{ .Name }}
├── Decl = {{ print .Decl }}
├── Data = {{ print .Data }}
└── Type = {{ print .Type }}{{ println }}`,
		},
	}
)
var (
	identTemplate  *template.Template
	objectTemplate *template.Template
)

func TemplNode(w io.Writer, fs *token.FileSet, node ast.Node) error {
	err := Initialize(fs)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("initialize"))
	}
	err = tmpltree(w, "", []string{"", ""}, node)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("execute tree"))
	}
	return nil
}

func Initialize(fs *token.FileSet) error {
	for k, t := range tmplMap {
		t := t
		v := strings.Split(t.Tpl, "\n")
		v = append([]string{""}, v...)
		s := "{{ template \"prefix1\" }}" + v[0] + strings.Join(v[1:], "\n"+"{{ template \"prefix2\" }}")
		funcMap := template.FuncMap{
			"tree": func(node ast.Node) string {
				return ""
			},
			"position": func(p token.Pos) token.Position {
				return fs.Position(p)
			},
			"obj": func(obj *ast.Object) string {
				return ""
			},
		}

		tmpl, err := template.New(k).Funcs(funcMap).Parse(s)
		if err != nil {
			return failure.Wrap(err, failure.Messagef("parse template %s", k))
		}
		t.Template = tmpl
		tmplMap[k] = t
	}
	return nil
}

func tmplident(w io.Writer, parentPrefix string, prefixes []string, node *ast.Ident) error {
	tpl := tmplMap[fmt.Sprintf("%T", node)]
	if tpl.Template == nil {
		return errors.New("Ident template is nil")
	}
	t, err := tpl.Template.Clone()
	if err != nil {
		return failure.Wrap(err, failure.Messagef("clone base ident template"))
	}
	funcMap := template.FuncMap{
		"tree": func(node ast.Node) string {
			tmpltree(w, parentPrefix+prefixes[0], middlePrefixes, node)
			return ""
		},
		"obj": func(obj *ast.Object) string {
			err := tmplobject(w, parentPrefix+prefixes[1]+tailLine, tailPrefixes, obj)
			if err != nil {
				return err.Error()
			}
			return ""
		},
	}
	t, err = t.Funcs(funcMap).Parse(fmt.Sprintf("{{define \"prefix1\"}}%s{{end}}{{ define \"prefix2\" }}%s{{end}}", parentPrefix+prefixes[0], parentPrefix+prefixes[1]))
	if err != nil {
		return failure.Wrap(err, failure.Messagef("parse prefix templates"))
	}
	err = t.Execute(w, node)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("execute ident template"))
	}
	return nil
}

func tmplobject(w io.Writer, parentPrefix string, prefixes []string, node *ast.Object) error {
	tpl := tmplMap[fmt.Sprintf("%T", node)]
	if tpl.Template == nil {
		return errors.New("object template is nil")
	}
	t, err := tpl.Template.Clone()
	if err != nil {
		return failure.Wrap(err, failure.Messagef("clone base object template"))
	}
	t, err = t.Parse(fmt.Sprintf("{{define \"prefix1\"}}%s{{end}}{{ define \"prefix2\" }}%s{{end}}", parentPrefix+prefixes[0], parentPrefix+prefixes[1]))
	if err != nil {
		return failure.Wrap(err, failure.Messagef("parse prefix templates"))
	}
	err = t.Execute(w, node)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("execute object template"))
	}
	return nil
}

func tmpltree(w io.Writer, parentPrefix string, prefixes []string, node ast.Node) error {
	switch n := node.(type) {
	case *ast.Ident:
		err := tmplident(w, parentPrefix, prefixes, n)
		if err != nil {
			return failure.Wrap(err, failure.Messagef("execute tree ident template"))
		}
	}
	return nil
}

// テンプレート文字列に対して先頭に罫線用のプレフィックスをつける
// template を全部ブロックにする
// 関数を定義する
//
