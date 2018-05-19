package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
)

// Parser :
type Parser struct {
	fileSet *token.FileSet

	Structs map[string]*StructDecl
}

// Parse :
func (p *Parser) Parse(code string, name string) error {
	file, err := parser.ParseFile(p.fileSet, name, code, 0)
	if err != nil {
		return err
	}

	ci := NewCodeInfo(file, code, name)
	for name, spec := range GetGenStructDecls(file) {
		structDecl := NewStructDecl(ci, name)
		p.Structs[name] = structDecl
		for _, field := range spec.Fields.List {
			fieldType := GetExprType(field.Type)
			for _, name := range field.Names {
				structDecl.Fields[name.Name] = NewStructFieldDecl(ci, structDecl, name.Name, fieldType)
			}
		}
	}
	return nil
}

// ParseFile :
func (p *Parser) ParseFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return p.Parse(string(data), path)
}

// NewParser :
func NewParser() *Parser {
	return &Parser{
		fileSet: token.NewFileSet(),
		Structs: make(map[string]*StructDecl),
	}
}
