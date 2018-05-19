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
	Tags   map[string]string
}

// IsBase :
func (f *StructFieldDecl) IsBase() bool {
	return f.Name == ""
}

// GetOptions :
func (f *StructFieldDecl) GetOptions() map[string]string {
	options := make(map[string]string)
	if f.Tags != nil {
		for k, v := range f.Tags {
			options[k] = v
		}

		_, ok := f.Tags["*"]
		if ok {
			options["get"] = ""
			options["set"] = ""
		}
	}
	return options
}

// NewStructFieldDecl :
func NewStructFieldDecl(ci *CodeInfo, s *StructDecl, name string, typ string, tags map[string]string) *StructFieldDecl {
	return &StructFieldDecl{
		DeclType: NewDeclType(ci, name),
		Struct:   s,
		Type:     typ,
		Tags:     tags,
	}
}
