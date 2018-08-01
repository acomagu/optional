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
