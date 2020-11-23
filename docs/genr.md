# genr
--
    import "github.com/openacid/genr"

Package genr provides with utilities to generate codes It builds a user defined
type `typeName` with type argument `valueType`, to emulate a generic type.

## Usage

#### func  Render

```go
func Render(fn string, header string, tmpl string, datas []interface{}, linters []string)
```
Render generate a file "fn". File content is defined by a "header", a repeated
body template "tmpl" and slice of "data" to render the body template.
Additianlly some linters can be specified to run after generating. Supported
linters are "gofmt" and "unconvert".

#### type IntConfig

```go
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
```

IntConfig defines a integer type based template redner config.

#### func  NewIntConfig

```go
func NewIntConfig(typeName, valueType string) *IntConfig
```
NewIntConfig build a IntConfig for a user defined type `typeName` with type
argument valueType, to emulate a generic type.
