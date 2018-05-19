package main

import (
	"bytes"
	"fmt"
	"go/src/html/template"
	"strings"
)

var getterTemplateStr = `
// {{.FuncName}} return {{.FieldName}}
func ({{.Receiver}} *{{.StructName}}) {{.FuncName}}() {{.FieldType}} {
	return {{.Receiver}}.{{.FieldName}}
}
`

var setterTemplateStr = `
// {{.FuncName}} set {{.FieldName}}
func ({{.Receiver}} *{{.StructName}}) {{.FuncName}}(value {{.FieldType}}) {
	{{.Receiver}}.{{.FieldName}} = value
}
`

// PropertyGenerator :
type PropertyGenerator struct {
	StructName string
	FieldName  string
	Receiver   string
	FuncName   string
	FieldType  string
	template   *template.Template
}

// Generate :
func (p *PropertyGenerator) Generate() string {
	var buffer bytes.Buffer
	err := p.template.Execute(&buffer, p)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// NewPropertyGenerator :
func NewPropertyGenerator(f *StructFieldDecl, s *StructDecl, fn string, temp *template.Template) *PropertyGenerator {
	return &PropertyGenerator{
		FieldName:  f.Name,
		StructName: s.Name,
		Receiver:   strings.ToLower(GetFirstChar(s.Name)),
		FuncName:   fn,
		FieldType:  f.Type,
		template:   temp,
	}
}

// NewGetter :
func NewGetter(f *StructFieldDecl, s *StructDecl) *PropertyGenerator {
	temp, err := template.New("getter").Parse(getterTemplateStr)
	if err != nil {
		panic(err)
	}

	fn := strings.Title(f.Name)
	firstCahr := GetFirstChar(f.Name)
	if strings.ToUpper(firstCahr) == firstCahr {
		fn = fmt.Sprintf("Get%s", fn)
	}
	return NewPropertyGenerator(
		f, s, fn, temp,
	)
}

// NewSetter :
func NewSetter(f *StructFieldDecl, s *StructDecl) *PropertyGenerator {
	temp, err := template.New("setter").Parse(setterTemplateStr)
	if err != nil {
		panic(err)
	}

	return NewPropertyGenerator(
		f, s, fmt.Sprintf("Set%s", f.Name), temp,
	)
}
