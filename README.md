# genr

[![Travis](https://travis-ci.com/openacid/genr.svg?branch=main)](https://travis-ci.com/openacid/genr)
![test](https://github.com/openacid/genr/workflows/test/badge.svg)

[![Report card](https://goreportcard.com/badge/github.com/openacid/genr)](https://goreportcard.com/report/github.com/openacid/genr)
[![Coverage Status](https://coveralls.io/repos/github/openacid/genr/badge.svg?branch=main&service=github)](https://coveralls.io/github/openacid/genr?branch=main&service=github)

[![GoDoc](https://godoc.org/github.com/openacid/genr?status.svg)](http://godoc.org/github.com/openacid/genr)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/openacid/genr)](https://pkg.go.dev/github.com/openacid/genr)
[![Sourcegraph](https://sourcegraph.com/github.com/openacid/genr/-/badge.svg)](https://sourcegraph.com/github.com/openacid/genr?badge)

genr generates source code to emulate `generic type`.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Synopsis](#synopsis)
  - [Generate two types `U16` and `I64` with a template:](#generate-two-types-u16-and-i64-with-a-template)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


# Synopsis

## Generate two types `U16` and `I64` with a template:

```go
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

```

The generated codes looks like the following:

```go
// Code generated 'by go generate ./...'; DO NOT EDIT.

package intarray

import (
	"encoding/binary"
)

type U16 struct {
	vals    []uint16
	eltSize int
	first   []byte
}

func NewU16(elts []uint16) (a *U16, err error) {
	a = &U16{
		vals:    make([]uint16, 10),
		eltSize: 2,
		first:   make([]byte, 0),
	}

	binary.LittleEndian.PutUint16(a.first, elts[0])
	return a, nil
}

type I64 struct {
	vals    []int64
	eltSize int
	first   []byte
}

func NewI64(elts []int64) (a *I64, err error) {
	a = &I64{
		vals:    make([]int64, 10),
		eltSize: 8,
		first:   make([]byte, 0),
	}

	binary.LittleEndian.PutUint64(a.first, uint64(elts[0]))
	return a, nil
}
```