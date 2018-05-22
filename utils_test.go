package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

type T struct {
	value struct {
		v1, v2 int
		v3     string
		v4     *string
		string
	} "props:\"+\""
}

// SetValue set value
func (t *T) SetValue(value struct {
	v1, v2 int
	v3     string
	v4     *string
	string
}) {
	t.value = value
}

// Value return value
func (t *T) Value() struct {
	v1, v2 int
	v3     string
	v4     *string
	string
} {
	return t.value
}

func TestFuncType(t *testing.T) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", `
package simba

type T struct {
	value struct {
		v, b int
		string
	} "props:\"+\""
}
		`, 0)
	ast.Print(fset, f)
}
