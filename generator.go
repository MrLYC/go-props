package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// Generator :
type Generator interface {
	Generate() string
}

// NotingGenerator :
type NotingGenerator struct{}

// Generate :
func (g *NotingGenerator) Generate() string {
	return ""
}

// Property :
type Property struct {
	Struct *StructDecl
	Field  *StructFieldDecl
	Getter Generator
	Setter Generator
}

// NewProperty :
func NewProperty(f *StructFieldDecl, s *StructDecl) *Property {
	return &Property{
		Struct: s,
		Field:  f,
		Getter: NewGetter(f, s),
		Setter: NewSetter(f, s),
	}
}

// PropertyManager :
type PropertyManager struct {
	Parser     *Parser
	Properties []*Property
}

// Len :
func (p *PropertyManager) Len() int {
	return len(p.Properties)
}

// Swap :
func (p *PropertyManager) Swap(i int, j int) {
	p.Properties[i], p.Properties[j] = p.Properties[j], p.Properties[i]
}

// Less :
func (p *PropertyManager) Less(i int, j int) bool {
	pi := p.Properties[i]
	pj := p.Properties[j]

	val := strings.Compare(pi.Struct.Name, pj.Struct.Name)
	if val == 0 {
		val = strings.Compare(pi.Field.Name, pj.Field.Name)
	}
	return val < 0
}

// Sort :
func (p *PropertyManager) Sort() {
	sort.Sort(p)
}

// Generate :
func (p *PropertyManager) Generate() string {
	var code string
	codeList := make([]string, 0)
	if Config.Package != "" {
		codeList = append(codeList, fmt.Sprintf("package %v\n", Config.Package))
	}
	p.Sort()
	for _, property := range p.Properties {
		codes := []string{
			property.Setter.Generate(),
			property.Getter.Generate(),
		}
		for _, code = range codes {
			if code != "" {
				codeList = append(codeList, code)
			}
		}
	}
	code = strings.Join(codeList, Config.LineSep)

	return code
}

// NewPropertyManager :
func NewPropertyManager(parser *Parser) Generator {
	properties := make([]*Property, 0)
	for _, s := range parser.Structs {
		if Config.WithoutPrivateStruct && !s.IsPublic() {
			log.Printf("ignore private struct: %v", s.Name)
			continue
		}
		for _, f := range s.Fields {
			if Config.TagName != "" && !f.Options.IsValid() {
				log.Printf("ignore field with invalid options: %v", f.Name)
				continue
			}
			if Config.WithoutPublicField && f.IsPublic() {
				log.Printf("ignore public field: %v", f.Name)
				continue
			}
			property := NewProperty(f, s)
			properties = append(properties, property)
		}
	}
	return &PropertyManager{
		Parser:     parser,
		Properties: properties,
	}
}
