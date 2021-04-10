package astree

// generation code
// https://play.golang.org/p/P2wzZ1rFRnU

import (
	_ "embed"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"text/template"

	"github.com/morikuni/failure"
)

var (
	//go:embed all.tpl
	tpl          string
	allTemplates *template.Template
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

	tmpl, err := template.New("all").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("parse template"))
	}
	allTemplates = tmpl
	return nil
}

func tmplobject(w io.Writer, parentPrefix string, prefixes []string, node *ast.Object) error {
	t, err := allTemplates.Clone()
	if err != nil {
		return failure.Wrap(err, failure.Messagef("clone base object template"))
	}
	t, err = t.Parse(fmt.Sprintf("{{define \"prefix1\"}}%s{{end}}{{ define \"prefix2\" }}%s{{end}}", parentPrefix+prefixes[0], parentPrefix+prefixes[1]))
	if err != nil {
		return failure.Wrap(err, failure.Messagef("parse prefix templates"))
	}
	err = t.ExecuteTemplate(w, fmt.Sprintf("%T", node), node)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("execute ident template"))
	}
	return nil
}

func tmpltree(w io.Writer, parentPrefix string, prefixes []string, node ast.Node) error {
	t, err := allTemplates.Clone()
	if err != nil {
		return failure.Wrap(err, failure.Messagef("clone base template"))
	}
	funcMap := template.FuncMap{
		"tree": func(node ast.Node) string {
			tmpltree(w, parentPrefix+prefixes[1]+middleLine, middlePrefixes, node)
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
	err = t.ExecuteTemplate(w, fmt.Sprintf("%T", node), node)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("execute %T template", node))
	}
	return nil
}
