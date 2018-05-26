package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"strconv"
)

// Parser :
type Parser struct {
	fileSet *token.FileSet

	Structs map[string]*StructDecl
}

func (p *Parser) parseStructDecl(ci *CodeInfo, structName string, spec *ast.StructType) error {
	structDecl := NewStructDecl(ci, structName)
	p.Structs[structName] = structDecl
	for _, field := range spec.Fields.List {
		var tags map[string]string
		if field.Tag != nil {
			tag, err := strconv.Unquote(field.Tag.Value)
			if err != nil {
				return err
			}
			tags = ParseTags(Config.TagName, tag)
		}
		fieldType := GetExprTypeLit(field.Type)

		for _, field := range field.Names {
			log.Printf("scaning struct %v field: %v %v[%+v]", structName, field.Name, fieldType, tags)
			structDecl.Fields[field.Name] = NewStructFieldDecl(ci, structDecl, field.Name, fieldType, tags)
		}
	}
	return nil
}
func (p *Parser) parseStructMethods(ci *CodeInfo, structName string, specs []*ast.FuncDecl) error {
	structDecl, ok := p.Structs[structName]
	if !ok {
		return ErrNotFound
	}
	for _, spec := range specs {
		structDecl.Methods[spec.Name.Name] = NewStructFuncDecl(ci, structDecl, spec.Name.Name)
	}
	return nil
}

// Parse :
func (p *Parser) Parse(code string, name string) error {
	file, err := parser.ParseFile(p.fileSet, name, code, 0)
	if err != nil {
		return err
	}

	ci := NewCodeInfo(file, code, name)
	for structName, structDecl := range GetGenStructDecls(file) {
		log.Printf("scaning struct %v fields", structName)
		err = p.parseStructDecl(ci, structName, structDecl)
		if err != nil {
			return err
		}
	}

	for structName, funcDecl := range GetStructMethodDecls(file) {
		log.Printf("scaning struct %v methods", structName)
		err = p.parseStructMethods(ci, structName, funcDecl)
		if err == ErrNotFound {
			log.Printf("unknown receiver %v", structName)
			continue
		} else if err != nil {
			return err
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
func (p *Parser) Prepare() *PropertyManager {
	return NewPropertyManager(p)
}

// NewParser :
func NewParser() *Parser {
	return &Parser{
		fileSet: token.NewFileSet(),
		Structs: make(map[string]*StructDecl),
	}
}
