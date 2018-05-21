package main

import (
	"fmt"
	"go/ast"
	"log"
	"reflect"
	"strings"
)

// Assert :
func Assert(value bool) {
	if !value {
		panic(ErrAssert)
	}
}

// GetChanTypeLit :
func GetChanTypeLit(t *ast.ChanType) string {
	chType := GetExprTypeLit(t.Value)
	switch t.Dir {
	case ast.SEND:
		return fmt.Sprintf("chan<- %s", chType)
	case ast.RECV:
		return fmt.Sprintf("<-chan %s", chType)
	}
	return fmt.Sprintf("chan %s", chType)
}

// GetFuncTypeLit :
func GetFuncTypeLit(t *ast.FuncType) string {
	params := make([]string, 0)
	if t.Params != nil {
		for _, p := range t.Params.List {
			params = append(params, GetExprTypeLit(p.Type))
		}
	}

	results := make([]string, 0)
	if t.Results != nil {
		for _, r := range t.Results.List {
			results = append(results, GetExprTypeLit(r.Type))
		}
	}

	paramsLit := strings.Join(params, ", ")
	resultsLit := strings.Join(results, ", ")
	switch len(results) {
	case 0:
		return fmt.Sprintf("func (%v)", paramsLit)
	case 1:
		return fmt.Sprintf("func (%v) %v", paramsLit, resultsLit)
	default:
		return fmt.Sprintf("func (%v) (%v)", paramsLit, resultsLit)
	}
}

// GetExprTypeLit :
func GetExprTypeLit(expr ast.Expr) string {
	switch typ := expr.(type) {
	case *ast.Ident:
		return typ.Name
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", GetExprTypeLit(typ.X))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", typ.X, typ.Sel)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", GetExprTypeLit(typ.Key), GetExprTypeLit(typ.Value))
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", GetExprTypeLit(typ.Elt))
	case *ast.ChanType:
		return GetChanTypeLit(typ)
	case *ast.InterfaceType:
		return fmt.Sprintf("interface{}")
	case *ast.FuncType:
		return GetFuncTypeLit(typ)
	}
	panic(fmt.Errorf("parse expr type failed at %v: %s%+v", expr.Pos(), reflect.TypeOf(expr).String(), expr))
}

// GetFuncDecls :
func GetFuncDecls(astFile *ast.File) []*ast.FuncDecl {
	decls := make([]*ast.FuncDecl, 0)
	for _, d := range astFile.Decls {
		switch decl := d.(type) {
		case *ast.FuncDecl:
			decls = append(decls, decl)
		}
	}
	return decls
}

// GetGenDecls :
func GetGenDecls(astFile *ast.File) []*ast.GenDecl {
	decls := make([]*ast.GenDecl, 0)
	for _, d := range astFile.Decls {
		switch decl := d.(type) {
		case *ast.GenDecl:
			decls = append(decls, decl)
		}
	}
	return decls
}

// GetGenStructDecls :
func GetGenStructDecls(astFile *ast.File) map[string]*ast.StructType {
	decls := make(map[string]*ast.StructType)
	for _, d := range GetGenDecls(astFile) {
		for _, s := range d.Specs {
			switch spec := s.(type) {
			case *ast.TypeSpec:
				specType, ok := spec.Type.(*ast.StructType)
				if ok {
					decls[spec.Name.Name] = specType
				}
			}
		}
	}
	return decls
}

// GetStructMethodDecls :
func GetStructMethodDecls(astFile *ast.File) map[string][]*ast.FuncDecl {
	decls := make(map[string][]*ast.FuncDecl)
	for _, d := range GetFuncDecls(astFile) {
		if d.Recv == nil {
			continue
		}
		for _, r := range d.Recv.List {
			typ := GetExprTypeLit(r.Type)
			if GetFirstChar(typ) == "*" {
				typ = typ[1:]
			}
			methods, ok := decls[typ]
			if !ok {
				methods = make([]*ast.FuncDecl, 0)
			}
			decls[typ] = append(methods, d)
		}
	}
	return decls
}

// GetFirstChar :
func GetFirstChar(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}
	return string(r[0])
}

// ParseTags :
func ParseTags(key string, content string) map[string]string {
	if key == "" {
		return nil
	}
	values := make(map[string]string)
	tag := reflect.StructTag(content)
	value, ok := tag.Lookup(key)
	if !ok {
		log.Printf("lookup key %v in tag failed: %v", key, tag)
		return nil
	}
	for _, opt := range strings.Split(value, ",") {
		parts := strings.SplitN(opt, "=", 2)
		switch len(parts) {
		case 1:
			values[strings.TrimSpace(parts[0])] = ""
		case 2:
			values[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return values
}
