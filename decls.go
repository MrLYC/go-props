package main

// StructDecl :
type StructDecl struct {
	*DeclType
	Fields map[string]*StructFieldDecl
}

// NewStructDecl :
func NewStructDecl(ci *CodeInfo, name string) *StructDecl {
	return &StructDecl{
		DeclType: NewDeclType(ci, name),
		Fields:   make(map[string]*StructFieldDecl),
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
