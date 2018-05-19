package main

import (
	"log"
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
	Getter Generator
	Setter Generator
}

// NewProperty :
func NewProperty(f *StructFieldDecl, s *StructDecl, options map[string]string) *Property {
	return &Property{
		Getter: NewGetter(f, s, options),
		Setter: NewSetter(f, s, options),
	}
}

// PropertyManager :
type PropertyManager struct {
	Parser     *Parser
	Properties []*Property
}

// Generate :
func (p *PropertyManager) Generate() string {
	codeList := make([]string, 0)
	for _, property := range p.Properties {
		codeList = append(codeList, property.Setter.Generate())
		codeList = append(codeList, property.Getter.Generate())
	}
	return strings.Join(codeList, Config.LineSep)
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
			if Config.TagName != "" && f.Tags == nil {
				log.Printf("ignore field without tag: %v", f.Name)
				continue
			}
			if Config.WithoutPublicField && f.IsPublic() {
				log.Printf("ignore public field: %v", f.Name)
				continue
			}
			properties = append(properties, NewProperty(f, s, f.GetOptions()))
		}
	}
	return &PropertyManager{
		Parser:     parser,
		Properties: properties,
	}
}
