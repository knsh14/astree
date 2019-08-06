package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/knsh14/astree"
	"github.com/morikuni/failure"
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
		os.Exit(1)
	}
	stat, err := os.Stat(abs)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fs := token.NewFileSet()
	if stat.IsDir() {
		err = packages(fs, args[0])
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	err = file(fs, args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func packages(fs *token.FileSet, p string) error {
	pkgs, err := parser.ParseDir(fs, p, func(fi os.FileInfo) bool {
		return true
	}, 0)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("failed to parse directory:", p))
	}
	astree.Packages(os.Stdout, pkgs)
	return nil
}

func file(fs *token.FileSet, p string) error {
	f, err := parser.ParseFile(fs, p, nil, 0)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("failed to parse file:", p))
	}
	astree.File(os.Stdout, f)
	return nil
}
