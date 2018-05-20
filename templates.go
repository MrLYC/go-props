package main

import (
	"bytes"
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
	Assert(p.StructName != "")
	Assert(p.FieldName != "")
	Assert(p.Receiver != "")
	Assert(p.FuncName != "")
	Assert(p.FieldType != "")

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
func NewGetter(f *StructFieldDecl, s *StructDecl) Generator {
	temp, err := template.New("getter").Parse(getterTemplateStr)
	if err != nil {
		panic(err)
	}

	fn, ok := f.Options.Get(GetterTag)
	if !ok || fn == DisableTag {
		log.Printf("getter not set: %v", f.Name)
		return &NotingGenerator{}
	}
	if f.IsGetterFound() {
		log.Printf("%v.%v getter has been declared", f.Struct.Name, f.Name)
		return &NotingGenerator{}
	}
	return NewPropertyGenerator(
		f, s, fn, temp,
	)
}

// NewSetter :
func NewSetter(f *StructFieldDecl, s *StructDecl) Generator {
	temp, err := template.New("setter").Parse(setterTemplateStr)
	if err != nil {
		panic(err)
	}
	fn, ok := f.Options.Get(SetterTag)
	if !ok || fn == DisableTag {
		log.Printf("setter not set: %v", f.Name)
		return &NotingGenerator{}
	}
	if f.IsSetterFound() {
		log.Printf("%v.%v setter has been declared", f.Struct.Name, f.Name)
		return &NotingGenerator{}
	}
	return NewPropertyGenerator(
		f, s, fn, temp,
	)
}
