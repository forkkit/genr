package genr

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var implHead = `package genr
import (
	_ "math/bits"
	"encoding/binary"
)
`

var implTemplate = `
type {{.Name}} struct {
	Base
}

func New{{.Name}}(elts []{{.ValType}}) (a *{{.Name}}, err error) {
	a := &{{.Name}}{}
	var size int  = {{.ValLen}}
	var v {{.EncodeCast}}
	buf := []byte{}
	binary.LittleEndian.Put{{.Codec}}(buf, {{.EncodeCast}}(elts[0]))
	return a,              nil
}

`

func TestRender(t *testing.T) {

	ta := require.New(t)

	pref := "int"
	implfn := pref + ".go"

	impls := []interface{}{
		NewIntConfig("U16", "uint16"),
		NewIntConfig("I64", "int64"),
	}

	{

		Render(implfn, implHead, implTemplate, impls, []string{})
		defer os.Remove(implfn)

		want := `
// Code generated 'by go generate ./...'; DO NOT EDIT.

package genr
import (
	_ "math/bits"
	"encoding/binary"
)


type U16 struct {
	Base
}

func NewU16(elts []uint16) (a *U16, err error) {
	a := &U16{}
	var size int  = 2
	var v uint16
	buf := []byte{}
	binary.LittleEndian.PutUint16(buf, uint16(elts[0]))
	return a,              nil
}


type I64 struct {
	Base
}

func NewI64(elts []int64) (a *I64, err error) {
	a := &I64{}
	var size int  = 8
	var v uint64
	buf := []byte{}
	binary.LittleEndian.PutUint64(buf, uint64(elts[0]))
	return a,              nil
}

`[1:]

		dat, err := ioutil.ReadFile(implfn)
		fmt.Println(err)
		ta.Nil(err)

		s := string(dat)
		ta.Equal(want, s)
	}

	{

		implfn := pref + ".go"
		Render(implfn, implHead, implTemplate, impls, []string{"gofmt", "unconvert"})
		defer os.Remove(implfn)

		want := `
// Code generated 'by go generate ./...'; DO NOT EDIT.

package genr

import (
	"encoding/binary"
	_ "math/bits"
)

type U16 struct {
	Base
}

func NewU16(elts []uint16) (a *U16, err error) {
	a := &U16{}
	var size int = 2
	var v uint16
	buf := []byte{}
	binary.LittleEndian.PutUint16(buf, elts[0])
	return a, nil
}

type I64 struct {
	Base
}

func NewI64(elts []int64) (a *I64, err error) {
	a := &I64{}
	var size int = 8
	var v uint64
	buf := []byte{}
	binary.LittleEndian.PutUint64(buf, uint64(elts[0]))
	return a, nil
}
`[1:]

		dat, err := ioutil.ReadFile(implfn)
		fmt.Println(err)
		ta.Nil(err)

		s := string(dat)
		ta.Equal(want, s)

	}
}
