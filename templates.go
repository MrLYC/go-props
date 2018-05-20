package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
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
func NewGetter(f *StructFieldDecl, s *StructDecl, options map[string]string) Generator {
	temp, err := template.New("getter").Parse(getterTemplateStr)
	if err != nil {
		panic(err)
	}

	fn, ok := options["get"]
	if !ok {
		return &NotingGenerator{}
	}
	if fn == "" {
		if f.IsPublic() {
			fn = fmt.Sprintf("Get%s", f.Name)
		} else {
			fn = strings.Title(f.Name)
		}
	}
	return NewPropertyGenerator(
		f, s, fn, temp,
	)
}

// NewSetter :
func NewSetter(f *StructFieldDecl, s *StructDecl, options map[string]string) Generator {
	temp, err := template.New("setter").Parse(setterTemplateStr)
	if err != nil {
		panic(err)
	}
	fn, ok := options["set"]
	if !ok {
		log.Printf("setter not found: %v", f.Name)
		return &NotingGenerator{}
	}
	if fn == "" {
		fn = fmt.Sprintf("Set%s", strings.Title(f.Name))
	}

	return NewPropertyGenerator(
		f, s, fn, temp,
	)
}
