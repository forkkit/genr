package main

import (
	"github.com/openacid/genr"
)

var implHead = `package intarray
import (
	"encoding/binary"
)
`

var implTemplate = `
type {{.Name}} struct {
	vals []{{.ValType}}
	eltSize int
	first []byte
}

func New{{.Name}}(elts []{{.ValType}}) (a *{{.Name}}, err error) {
	a = &{{.Name}}{
		vals: make([]{{.ValType}}, 10),
		eltSize: {{.ValLen}},
		first: make([]byte, 0),
	}

	binary.LittleEndian.Put{{.Codec}}(a.first, {{.EncodeCast}}(elts[0]))
	return a, nil
}
`

func main() {

	implfn := "../intarray.go"

	impls := []interface{}{
		genr.NewIntConfig("U16", "uint16"),
		genr.NewIntConfig("I64", "int64"),
	}

	genr.Render(implfn, implHead, implTemplate, impls, []string{"gofmt", "unconvert"})

}
