//Package schema ...
//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
package schema

import (
	"bytes"
	"os"
)

// String : Run `go generate` against this package
func String(saveFile bool) string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}
	str := buf.String()
	if saveFile {
		SaveSchema(str)
	}
	return str
}

// SaveSchema : save the schema in a file
func SaveSchema(str string) {
	f, err := os.Create("./schema/schema.graphql.debug")
	if err != nil {
		panic("cant write schema log file" + err.Error())
	}
	if _, err := f.WriteString(str); err != nil {
		panic("cant flush string schema in log file" + err.Error())
	}
	if err := f.Sync(); err != nil {
		panic("cant sync schema log file" + err.Error())
	}
}
