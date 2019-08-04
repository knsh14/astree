package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/knsh14/astree"
)

var (
	ignorePattern string
)

func init() {
	flag.StringVar(&ignorePattern, "I", "", "")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "argument must be 1")
		os.Exit(2)
	}
	abs, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	stat, err := os.Stat(abs)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	fs := token.NewFileSet()
	if stat.IsDir() {
		return
	}
	f, err := parser.ParseFile(fs, abs, nil, 0)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	astree.Tree(os.Stdout, "", []string{"", ""}, f)

}
