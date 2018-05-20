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
package main

type T struct {
	f func(int)(T, error)
}

func (fffff *T)Test(value int) (data T, error) {

}
		`, 0)
	ast.Print(fset, f)
}
