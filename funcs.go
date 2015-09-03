package schematype

import (
	"bytes"
	"fmt"
	"strings"
)

const maxLineLength = 80

// TODO wrap at line length
func comment(value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("// %s", value)
}

func typeName(name string) string {
	return titleize(name)
}

func fieldName(name string) string {
	return titleize(name)
}

func fieldTag(name string, required bool) string {
	buf := bytes.NewBufferString("`json:\"")
	buf.WriteString(name)
	if !required {
		buf.WriteString(",omitempty")
	}
	buf.WriteString("\"`")
	return buf.String()
}

func titleize(value string) string {
	words := strings.Split(value, "_")
	parts := make([]string, len(words))
	for i, word := range words {
		parts[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(parts, "")
}
