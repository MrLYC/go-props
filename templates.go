package main

var getterTemplateStr = `
// {{.FuncName}} return {{.Field.Name}}
func ({{.Receiver}} *{{.Struct.Name}}) {{.FuncName}}() {{.Field.Type}} {
	return {{.Receiver}}.{{.Field.Name}}
}
`

var setterTemplateStr = `
// {{.FuncName}} set {{.Field.Name}}
func ({{.Receiver}} *{{.Struct.Name}}) {{.FuncName}}(v {{.Field.Type}}) {
	{{.Receiver}}.{{.Field.Name}} = v
}
`
