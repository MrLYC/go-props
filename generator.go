package main

import (
	"fmt"
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
	var code string
	codeList := make([]string, 0)
	if Config.Package != "" {
		codeList = append(codeList, fmt.Sprintf("package %v\n", Config.Package))
	}
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
