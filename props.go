package main

import (
	"go/ast"
)

// CodeInfo :
type CodeInfo struct {
	AST  *ast.File
	Code string
	Name string
}

// NewCodeInfo :
func NewCodeInfo(ast *ast.File, code string, name string) *CodeInfo {
	return &CodeInfo{
		AST:  ast,
		Code: code,
		Name: name,
	}
}

// DeclType :
type DeclType struct {
	*CodeInfo
}

// NewDeclType :
func NewDeclType(ci *CodeInfo) *DeclType {
	return &DeclType{
		CodeInfo: ci,
	}
}
