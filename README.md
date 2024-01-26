# A Tour of Go

**Current step:** https://go.dev/tour/basics/12

## Questions/Doubts

- CLI: Go watch mode? -> nope or https://github.com/mitranim/gow

## Basics

- **Go package:** folder that contains a collection of Go files
- **Go module:** collection of Go packages

### CLI

```shell
# initialize new project -> create new module
go mod init [module name]
go mod init [URL of the GitHub repository]

# compile
go build [filename]

# compile and run
go run [filename]

# test (only files ending with "_test.go")
go test [file or folder] -v
```

### Runtime

- Programs start running in package `main`

### Code

### Packages

- Exports: names are exported only if they begin with a capital letter

### Syntax

#### Variables

- The type of a variable comes after the variable name
- `var` statements can be at package or function level
- short assignment statements (e.g. `val := 7`):
  - can be used in place of a `var` declaration with implicit type
  - are only permitted inside a function

#### Types

**Basic**

- `bool`
- `string`
- `int` `int8` `int16` `int32` `int64`
- `uint` `uint8` `uint16` `uint32` `uint64` `uintptr`
- `byte` // alias for uint8
- `rune` // alias for int32, represents a Unicode code point
- `float32` `float64`
- `complex64` `complex128`

#### Zero values

Variables declared without an explicit initial value are given their **zero value**:

- `0` for numeric types,
- `false` for the boolean type
- `""` (empty string) for strings

#### Functions

- A function can return any number of results!
- Return values may be named. If so, they are treated as variables defined at the top of the function
- "naked" return statement: `return` statement without arguments -> it returns the named return values (careful about readability -> only use in short functions)

### Testing (WiP)

- Files ending in `_test.go` are skipped
- Looks for function starting with `Test`, e.g.: `func TestIndexHandlerPrefix(t *testing.T) {}`
- Table-driven test with `t.Run()`

## Resources

- "Learn GO Fast (Full Tutorial)": https://www.youtube.com/watch?v=8uiZC0l4Ajw
- "A Tour of Go": https://go.dev/tour/basics/9
