package main

import (
	"fmt"
	"strings"
)

// PropsOptions :
type PropsOptions struct {
	isValid bool
	field   *StructFieldDecl
	Tags    map[string]string
}

// Get :
func (p *PropsOptions) Get(key string) (string, bool) {
	if p.IsValid() {
		val, ok := p.Tags[key]
		return val, ok
	}
	return "", false
}

// IsValid :
func (p *PropsOptions) IsValid() bool {
	return p.isValid
}

// Init :
func (p *PropsOptions) Init() {
	_, ok := p.Tags["*"]
	if ok {
		p.Tags["get"] = ""
		p.Tags["set"] = ""
	}
	p.initGetOpt()
	p.initSetOpt()
}
func (p *PropsOptions) initSetOpt() {
	if p.Tags["set"] == "" {
		p.Tags["set"] = fmt.Sprintf("Set%s", strings.Title(p.field.Name))
	}
}

func (p *PropsOptions) initGetOpt() {
	if p.Tags["get"] == "" {
		if p.field.IsPublic() {
			p.Tags["get"] = fmt.Sprintf("Get%s", p.field.Name)
		} else {
			p.Tags["get"] = strings.Title(p.field.Name)
		}
	}
}

// NewPropsOptions :
func NewPropsOptions(f *StructFieldDecl, tags map[string]string) *PropsOptions {
	opt := &PropsOptions{
		isValid: false,
		field:   f,
		Tags:    make(map[string]string),
	}
	if tags != nil {
		opt.isValid = true
		for k, v := range tags {
			opt.Tags[k] = v
		}
	}
	if opt.IsValid() {
		opt.Init()
	}
	return opt
}
