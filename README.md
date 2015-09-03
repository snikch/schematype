# schematype

Automatically generate Go types from JSON schema.

This is beta software, and the API *will* change.

# CLI

A cli exists in `cmd/schematype`, install via `go install github.com/snikch/schematype/cmd/schematype`.

```
cat examples/schema.json | schematype | gofmt
```

```go
package main

type MyType struct {
	Description *string `json:"description,omitempty"` // The description
	Name        string  `json:"name"`                  // The name
}
```

## Setting Type and Package names

Set the type and package with flags.

```
cat examples/schema.json | schematype --name SomeType --package types | gofmt
```

```go
package types

type SomeType struct {
	Description *string `json:"description,omitempty"` // The description
	Name        string  `json:"name"`                  // The name
}
```

## Multiple files

Read all schema json files in `types/schemas/*.json` and output them to `types/`. Each schema will get its own `.go` file with the same name as the `.json` schema file.

```
for filename in types/schemas/*.json; cat $filename | schematype --name=`basename $filename .json` --package=types | gofmt > types/`basename $filename .json`.go
```

