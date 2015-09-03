# schematype

Automatically generate Go types from JSON schema.

This is beta software, and the API *will* change.

# CLI

A cli exists in `cmd/schematype`, install via `go install github.com/snikch/schematype/cmd/schematype`.

```
cat my_schema.json | schematype | gofmt
```

```go
package main

type MyType struct {
	Description *string `json:"description,omitempty"` // The description
	Name        string  `json:"name"`                  // The name
}
```

Set the type and name with flags.

```
cat my_schema.json | schematype --name SomeType --package types | gofmt
```

```go
package types

type SomeType struct {
	Description *string `json:"description,omitempty"` // The description
	Name        string  `json:"name"`                  // The name
}
```
