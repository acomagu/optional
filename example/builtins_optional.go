// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

// String
type OptionalString struct {
	v   string
	has bool
}

func NoneString() OptionalString {
	return OptionalString{}
}

func SomeString(v string) OptionalString {
	return OptionalString{
		v:   v,
		has: true,
	}
}

func (o OptionalString) Get() (string, bool) {
	return o.v, o.has
}

// Int
type OptionalInt struct {
	v   int
	has bool
}

func NoneInt() OptionalInt {
	return OptionalInt{}
}

func SomeInt(v int) OptionalInt {
	return OptionalInt{
		v:   v,
		has: true,
	}
}

func (o OptionalInt) Get() (int, bool) {
	return o.v, o.has
}
