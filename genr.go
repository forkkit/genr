// Package genr provides with utilities to generate codes
// It builds a user defined type `typeName` with type
// argument `valueType`, to emulate a generic type.
package genr

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

// IntConfig defines a integer type based template redner config.
type IntConfig struct {
	// Name of the data type.
	Name string
	// ValType is the actual underlying int type, such as int32.
	ValType string
	// ValLen specifies the length of ValType
	ValLen int
	// Codec defines the name of function to decode raw bytes into ValType.
	Codec string
	// EncodeCast defines a cast type/function to convert values before encode.
	// Because sometimes encoder does not provides a exact type.
	EncodeCast string
}

var (
	valLenMap = map[string]int{
		"uint8":  1,
		"uint16": 2,
		"uint32": 4,
		"uint64": 8,
		"int8":   1,
		"int16":  2,
		"int32":  4,
		"int64":  8,
	}

	decoderMap = map[string]string{
		"uint8":  "Uint8",
		"uint16": "Uint16",
		"uint32": "Uint32",
		"uint64": "Uint64",
		"int8":   "Uint8",
		"int16":  "Uint16",
		"int32":  "Uint32",
		"int64":  "Uint64",
	}
	encodeCastMap = map[string]string{
		"uint8":  "uint8",
		"uint16": "uint16",
		"uint32": "uint32",
		"uint64": "uint64",
		"int8":   "uint8",
		"int16":  "uint16",
		"int32":  "uint32",
		"int64":  "uint64",
	}
)

// NewIntConfig build a IntConfig for a user defined type `typeName` with type
// argument valueType, to emulate a generic type.
func NewIntConfig(typeName, valueType string) *IntConfig {

	valLen := valLenMap[valueType]
	codec := decoderMap[valueType]
	encodeCast := encodeCastMap[valueType]

	return &IntConfig{
		Name:       typeName,
		ValType:    valueType,
		ValLen:     valLen,
		Codec:      codec,
		EncodeCast: encodeCast,
	}
}

// Render generate a file "fn".
// File content is defined by a "header", a repeated body template
// "tmpl" and slice of "data" to render the body template.
// Additianlly some linters can be specified to run after generating.
// Supported linters are "gofmt" and "unconvert".
func Render(fn string, header string, tmpl string, datas []interface{}, linters []string) {

	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	err = f.Truncate(0)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(f, "// Code generated 'by go generate ./...'; DO NOT EDIT.")
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, header)

	t, err := template.New("foo").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	for _, d := range datas {
		err = t.Execute(f, d)
		if err != nil {
			panic(err)
		}
	}
	err = f.Sync()
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	for _, linter := range linters {
		var cmds []string
		switch linter {
		case "gofmt":
			cmds = []string{"gofmt", "-s", "-w", fn}
		case "unconvert":
			cmds = []string{"unconvert", "-v", "-apply", "./"}
		default:
			panic("unknown linter:" + linter)
		}

		out, err := exec.Command(cmds[0], cmds[1:]...).CombinedOutput()
		if err != nil {
			fmt.Println(cmds)
			fmt.Println(string(out))
			fmt.Println(err)
			panic(err)
		}
	}
}
