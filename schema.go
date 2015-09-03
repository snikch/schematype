package schematype

import (
	"bytes"
	"fmt"
	"sort"
)

type SchemaProperty struct {
	Type        string                    `json:"type"`
	Description string                    `json:"description,omitempty"`
	Properties  map[string]SchemaProperty `json:"properties,omitempty"`
	Required    []string                  `json:"required,omitempty"`
	Format      string                    `json:"format,omitempty"`
	// Items           *Schema     `json:"items,omitempty"`
}

type NamedProperty struct {
	SchemaProperty
	Name            string
	IsRequiredField bool
}

type Schema struct {
	NamedProperty
	Schema string `json:"$schema,omitempty"`
}

func (s Schema) TypeString(name string) (string, error) {
	buffer := bytes.NewBuffer(nil)
	err := templates.ExecuteTemplate(buffer, "struct", struct {
		Schema
		Name string
	}{
		Schema: s,
		Name:   name,
	})
	return buffer.String(), err
}

func (s SchemaProperty) Fields() []NamedProperty {
	// Create a sorted array of property names
	names := make([]string, len(s.Properties))
	i := 0
	for name, _ := range s.Properties {
		names[i] = name
		i++
	}

	requiredLookup := map[string]bool{}
	for _, name := range s.Required {
		requiredLookup[name] = true
	}

	sort.Strings(names)
	properties := make([]NamedProperty, len(s.Properties))
	for i, name := range names {
		properties[i] = NamedProperty{
			SchemaProperty:  s.Properties[name],
			Name:            name,
			IsRequiredField: requiredLookup[name],
		}
	}
	return properties
}

func (s NamedProperty) GoType() (string, error) {
	baseType := ""
	switch s.Type {
	case "any":
		baseType = "interface{}"
	case "boolean":
		baseType = "bool"
	case "integer":
		baseType = "int"
	case "number":
		baseType = "float64"
	case "object":
		buf := bytes.NewBufferString("struct {")
		for name, prop := range s.Properties {
			err := templates.ExecuteTemplate(buf, "field", NamedProperty{
				SchemaProperty: prop,
				Name:           name,
			})
			if err != nil {
				return "", err
			}
		}
		buf.WriteString("\n}")
		baseType = buf.String()
	case "string":
		switch s.Format {
		case "date-time":
			baseType = "time.Time"
		default:
			baseType = "string"
		}
	// case "array":
	// 	if s.Items != {
	// 		baseType = "[]" + s.Items.GoType()
	// 	} else {
	// 		baseType = "[]interface{}"
	// 	}
	case "null":
		fallthrough
	default:
		return "", fmt.Errorf("unknown type %s", s.Type)
	}

	if s.IsRequiredField {
		return baseType, nil
	}

	return "*" + baseType, nil
}
