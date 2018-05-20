package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"strings"
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
	for structName, spec := range GetGenStructDecls(file) {
		log.Printf("scaning struct %v", structName)
		structDecl := NewStructDecl(ci, structName)
		p.Structs[structName] = structDecl
		for _, field := range spec.Fields.List {
			var tags map[string]string
			if field.Tag != nil {
				tags = ParseTags(Config.TagName, strings.Trim(field.Tag.Value, "`"))
			}
			fieldType := GetExprTypeLit(field.Type)

			for _, field := range field.Names {
				log.Printf("scaning struct %v field: %v %v[%+v]", structName, field.Name, fieldType, tags)
				structDecl.Fields[field.Name] = NewStructFieldDecl(ci, structDecl, field.Name, fieldType, tags)
			}
		}
	}
	return nil
}

// ParseFile :
func (p *Parser) ParseFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("read file error: %v", err)
		return err
	}
	return p.Parse(string(data), path)
}

// Prepare :
func (p *Parser) Prepare() Generator {
	return NewPropertyManager(p)
}

// NewParser :
func NewParser() *Parser {
	return &Parser{
		fileSet: token.NewFileSet(),
		Structs: make(map[string]*StructDecl),
	}
}
