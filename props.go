package main

import (
	"go/ast"
	"strings"
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
	Name string
}

// IsPublic :
func (s *DeclType) IsPublic() bool {
	firstChar := GetFirstChar(s.Name)
	return strings.ToUpper(firstChar) == firstChar
}

// NewDeclType :
func NewDeclType(ci *CodeInfo, name string) *DeclType {
	return &DeclType{
		CodeInfo: ci,
		Name:     name,
	}
}
