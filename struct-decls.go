package main

// StructDecl :
type StructDecl struct {
	*DeclType
	Name   string
	Fields map[string]*StructFieldDecl
}

// NewStructDecl :
func NewStructDecl(ci *CodeInfo, name string) *StructDecl {
	return &StructDecl{
		DeclType: NewDeclType(ci),
		Name:     name,
		Fields:   make(map[string]*StructFieldDecl),
	}
}

// StructFieldDecl :
type StructFieldDecl struct {
	*DeclType
	Struct *StructDecl
	Name   string
	Type   string
}

// IsBase :
func (s *StructFieldDecl) IsBase() bool {
	return s.Name == ""
}

// NewStructFieldDecl :
func NewStructFieldDecl(ci *CodeInfo, s *StructDecl, name string, typ string) *StructFieldDecl {
	return &StructFieldDecl{
		DeclType: NewDeclType(ci),
		Struct:   s,
		Name:     name,
		Type:     typ,
	}
}
