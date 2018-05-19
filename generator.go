package main

import (
	"strings"
)

// Generator :
type Generator interface {
	Generate() string
}

// Property :
type Property struct {
	Getter Generator
	Setter Generator
}

// NewProperty :
func NewProperty(f *StructFieldDecl, s *StructDecl) *Property {
	return &Property{
		Getter: NewGetter(f, s),
		Setter: NewSetter(f, s),
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
		if !Config.WithPrivateStruct && s.IsPublic() {
			continue
		}
		for _, f := range s.Fields {
			if !Config.WithPublicField && f.IsPublic() {
				continue
			}
			properties = append(properties, NewProperty(f, s))
		}
	}
	return &PropertyManager{
		Parser:     parser,
		Properties: properties,
	}
}
