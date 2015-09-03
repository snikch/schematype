package schematype

import (
	"log"
	"text/template"
)

var templates *template.Template
var funcs = template.FuncMap{
	"comment":   comment,
	"typeName":  typeName,
	"fieldName": fieldName,
	"fieldTag":  fieldTag,
}

func init() {
	err := RegenerateTemplates()
	if err != nil {
		log.Fatal(err)
	}
}
func RegenerateTemplates() error {
	tmpl, err := template.New("schemaType").Funcs(funcs).Parse(Template)
	if err != nil {
		return err
	}
	templates = tmpl
	return nil
}

var (
	Template = `
{{define "file"}}
package {{.Package}}
{{end}}

{{define "struct"}}
  {{comment .Description}}
  type {{typeName .Name}} struct {
    {{range .Fields}}{{ template "field" .}}{{end}}
  }
{{end}}

{{define "field"}}
  {{fieldName .Name}} {{.GoType}} {{fieldTag .Name .IsRequiredField}} {{comment .Description}}{{end}}
`
)
