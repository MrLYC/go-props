package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFuncType(t *testing.T) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", `
package simba

type T struct {
	value map[*string]*int "props:\"+\""
}
		`, 0)
	ast.Print(fset, f)
}
