package main

import (
	"fmt"
	"log"
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
	_, ok := p.Tags[AutoSetterGetterTag]
	if ok {
		p.Tags[GetterTag] = ""
		p.Tags[SetterTag] = ""
	}
	p.initGetOpt()
	log.Printf("%v.%v getter: %v", p.field.Struct.Name, p.field.Name, p.Tags[GetterTag])
	p.initSetOpt()
	log.Printf("%v.%v setter: %v", p.field.Struct.Name, p.field.Name, p.Tags[SetterTag])
}
func (p *PropsOptions) initSetOpt() {
	if p.Tags[SetterTag] == "" {
		p.Tags[SetterTag] = fmt.Sprintf("Set%s", strings.Title(p.field.Name))
	}
}

func (p *PropsOptions) initGetOpt() {
	if p.Tags[GetterTag] == "" {
		if p.field.IsPublic() {
			p.Tags[GetterTag] = fmt.Sprintf("Get%s", p.field.Name)
		} else {
			p.Tags[GetterTag] = strings.Title(p.field.Name)
		}
	}
}

// NewPropsOptions :
func NewPropsOptions(f *StructFieldDecl, tags map[string]string) *PropsOptions {
	opt := &PropsOptions{
		isValid: false,
		field:   f,
		Tags: map[string]string{
			SetterTag: DisableTag,
			GetterTag: DisableTag,
		},
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
