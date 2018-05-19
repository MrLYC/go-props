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
	Struct *StructDecl
	Type   string
}

// IsBase :
func (s *StructFieldDecl) IsBase() bool {
	return s.Name == ""
}

// NewStructFieldDecl :
func NewStructFieldDecl(ci *CodeInfo, s *StructDecl, name string, typ string) *StructFieldDecl {
	return &StructFieldDecl{
		DeclType: NewDeclType(ci, name),
		Struct:   s,
		Type:     typ,
	}
}
