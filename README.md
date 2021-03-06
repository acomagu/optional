# optional

[![CircleCI](https://img.shields.io/circleci/project/github/acomagu/optional.svg?style=flat-square)](https://circleci.com/gh/acomagu/optional) ![GolangCI](https://golangci.com/badges/github.com/acomagu/optional.svg)

The optional type generator for Go.

## Example

```Go
package main

import "fmt"

//go:generate optional -type=string -name=String -output=builtins_optional.go
//go:generate optional -type=int -name=Int -output=builtins_optional.go -append

//go:generate optional -type=*T -name=T
type T struct{}

type Data struct {
	Int OptionalInt
	Str OptionalString
	T   OptionalT
}

func main() {
	d := Data{
		Int: SomeInt(3),
	}

	fmt.Println(d.Int.Get()) // 3, true
	fmt.Println(d.Str.Get()) // "", false
}
```

See [example/](./example) or type `optional --help` for more detail.
