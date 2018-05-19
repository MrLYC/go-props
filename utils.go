package main

import (
	"fmt"
	"go/ast"
	"reflect"
)

// GetExprType :
func GetExprType(expr ast.Expr) string {
	switch typ := expr.(type) {
	case *ast.Ident:
		return typ.Name
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", GetExprType(typ.X))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", typ.X, typ.Sel)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", typ.Key, GetExprType(typ.Value))
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", GetExprType(typ.Elt))
	case *ast.ChanType:
		return fmt.Sprintf("chan %s", GetExprType(typ.Value))
	}
	panic(fmt.Errorf("parse expr type failed at %v: %s%+v", expr.Pos(), reflect.TypeOf(expr).String(), expr))
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

// GetFirstChar :
func GetFirstChar(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}
	return string(r[0])
}
