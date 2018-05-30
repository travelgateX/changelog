package schema

import (
	"bytes"
	"io/ioutil"

	graphql "github.com/graph-gophers/graphql-go"
)

// ==========  <necessary explanation>  ==========
// ==========  <necessary explanation>  ==========
// ==========  <necessary explanation>  ==========
// Use `go generate` to pack all *.graphql files under this directory (and sub-directories) into
// a binary format.
//
//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...

// String reads the .graphql schema files from the generated _bindata.go file, concatenating the
// files together into one string.
//
// If this method complains about not finding functions AssetNames() or MustAsset(),
// run `go generate` against this package to generate the functions.
func String() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}

// MustParseSchema : load current graphql schema, panics on error
// this function loads schema from one unique file .grapqh
func MustParseSchema(queryResolver interface{}) *graphql.Schema {
	file, err := ioutil.ReadFile("./schema/schema.graphql")
	if err != nil {
		panic("cant read schema.graphql, " + err.Error())
	}

	schema, err := graphql.ParseSchema(string(file), queryResolver)
	if err != nil {
		panic("cant parse schema.graphql, " + err.Error())
	}

	return schema
}
