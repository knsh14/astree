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
	middlePrefix  string
	tailPrefix    string
	middleLine    string
	tailLine      string
)

func init() {
	flag.StringVar(&ignorePattern, "I", "", "")
	flag.StringVar(&middlePrefix, "middle-prefix", "├── ", "prefix for middle tree nodes")
	flag.StringVar(&tailPrefix, "tail-prefix", "└── ", "prefix for tail tree nodes")
	flag.StringVar(&middleLine, "middle-line", "│   ", "line continuation for middle nodes")
	flag.StringVar(&tailLine, "tail-line", "    ", "line continuation for tail nodes")
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
	config := astree.Config{
		MiddlePrefix: middlePrefix,
		TailPrefix:   tailPrefix,
		MiddleLine:   middleLine,
		TailLine:     tailLine,
	}
	if stat.IsDir() {
		err = packages(fs, args[0], config)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	err = file(fs, args[0], config)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func packages(fs *token.FileSet, p string, config astree.Config) error {
	pkgs, err := parser.ParseDir(fs, p, func(fi os.FileInfo) bool {
		return true
	}, 0)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("failed to parse directory: %s", p))
	}
	return astree.PackagesWithConfig(os.Stdout, fs, pkgs, config)
}

func file(fs *token.FileSet, p string, config astree.Config) error {
	f, err := parser.ParseFile(fs, p, nil, 0)
	if err != nil {
		return failure.Wrap(err, failure.Messagef("failed to parse file: %s", p))
	}
	return astree.FileWithConfig(os.Stdout, fs, f, config)
}
