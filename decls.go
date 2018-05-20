package main

// StructDecl :
type StructDecl struct {
	*DeclType
	Fields  map[string]*StructFieldDecl
	Methods map[string]*StructFuncDecl
}

// NewStructDecl :
func NewStructDecl(ci *CodeInfo, name string) *StructDecl {
	return &StructDecl{
		DeclType: NewDeclType(ci, name),
		Fields:   make(map[string]*StructFieldDecl),
		Methods:  make(map[string]*StructFuncDecl),
	}
}

// StructFuncDecl :
type StructFuncDecl struct {
	*DeclType
	Struct *StructDecl
}

// NewStructFuncDecl :
func NewStructFuncDecl(ci *CodeInfo, s *StructDecl, name string) *StructFuncDecl {
	return &StructFuncDecl{
		DeclType: NewDeclType(ci, name),
		Struct:   s,
	}
}

// StructFieldDecl :
type StructFieldDecl struct {
	*DeclType
	Struct  *StructDecl
	Type    string
	Options *PropsOptions
}

// IsBase :
func (f *StructFieldDecl) IsBase() bool {
	return f.Name == ""
}

func (f *StructFieldDecl) isMethodFound(tag string) bool {
	name, ok := f.Options.Get(tag)
	if !ok {
		return false
	}
	_, ok = f.Struct.Methods[name]
	return ok
}

// IsSetterFound :
func (f *StructFieldDecl) IsSetterFound() bool {
	return f.isMethodFound(SetterTag)
}

// IsGetterFound :
func (f *StructFieldDecl) IsGetterFound() bool {
	return f.isMethodFound(GetterTag)
}

// NewStructFieldDecl :
func NewStructFieldDecl(ci *CodeInfo, s *StructDecl, name string, typ string, tags map[string]string) *StructFieldDecl {
	f := &StructFieldDecl{
		DeclType: NewDeclType(ci, name),
		Struct:   s,
		Type:     typ,
	}
	f.Options = NewPropsOptions(f, tags)
	return f
}
