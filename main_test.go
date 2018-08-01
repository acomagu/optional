package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestGenerate(t *testing.T) {
	tests := map[string]struct {
		imports  []string
		pkgName  string
		typeName string
		optName  string
		append   bool
		before   string
		after    string
	}{
		"empty": {
			pkgName:  "main",
			typeName: "Auth",
			optName:  "Auth",
			append:   false,
			before:   "",
			after: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

// Auth
type OptionalAuth struct {
	v   Auth
	has bool
}

func NoneAuth() OptionalAuth {
	return OptionalAuth{}
}

func SomeAuth(v Auth) OptionalAuth {
	return OptionalAuth{
		v:   v,
		has: true,
	}
}

func (o OptionalAuth) Get() (Auth, bool) {
	return o.v, o.has
}
`,
		},
		"append": {
			append:   true,
			pkgName:  "main",
			typeName: "*Byte",
			optName:  "Byte",
			before: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

// Auth
type OptionalAuth struct {
	v   Auth
	has bool
}

func NoneAuth() OptionalAuth {
	return OptionalAuth{}
}

func SomeAuth(v Auth) OptionalAuth {
	return OptionalAuth{
		v:   v,
		has: true,
	}
}

func (o OptionalAuth) Get() (Auth, bool) {
	return o.v, o.has
}
`,
			after: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

// Auth
type OptionalAuth struct {
	v   Auth
	has bool
}

func NoneAuth() OptionalAuth {
	return OptionalAuth{}
}

func SomeAuth(v Auth) OptionalAuth {
	return OptionalAuth{
		v:   v,
		has: true,
	}
}

func (o OptionalAuth) Get() (Auth, bool) {
	return o.v, o.has
}

// Byte
type OptionalByte struct {
	v   *Byte
	has bool
}

func NoneByte() OptionalByte {
	return OptionalByte{}
}

func SomeByte(v *Byte) OptionalByte {
	return OptionalByte{
		v:   v,
		has: true,
	}
}

func (o OptionalByte) Get() (*Byte, bool) {
	return o.v, o.has
}
`,
		},
		"optName": {
			append:   false,
			pkgName:  "main",
			typeName: "*int",
			optName:  "Count",
			before:   "",
			after: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

// Count
type OptionalCount struct {
	v   *int
	has bool
}

func NoneCount() OptionalCount {
	return OptionalCount{}
}

func SomeCount(v *int) OptionalCount {
	return OptionalCount{
		v:   v,
		has: true,
	}
}

func (o OptionalCount) Get() (*int, bool) {
	return o.v, o.has
}
`,
		},
		"override": {
			pkgName:  "main",
			typeName: "*int",
			optName:  "Count",
			append:   false,
			before: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

// Auth
type OptionalAuth struct {
	v   Auth
	has bool
}

func NoneAuth() OptionalAuth {
	return OptionalAuth{}
}

func SomeAuth(v Auth) OptionalAuth {
	return OptionalAuth{
		v:   v,
		has: true,
	}
}

func (o OptionalAuth) Get() (Auth, bool) {
	return o.v, o.has
}

// Byte
type OptionalByte struct {
	v   *Byte
	has bool
}

func NoneByte() OptionalByte {
	return OptionalByte{}
}

func SomeByte(v *Byte) OptionalByte {
	return OptionalByte{
		v:   v,
		has: true,
	}
}

func (o OptionalByte) Get() (*Byte, bool) {
	return o.v, o.has
}
`,
			after: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

// Count
type OptionalCount struct {
	v   *int
	has bool
}

func NoneCount() OptionalCount {
	return OptionalCount{}
}

func SomeCount(v *int) OptionalCount {
	return OptionalCount{
		v:   v,
		has: true,
	}
}

func (o OptionalCount) Get() (*int, bool) {
	return o.v, o.has
}
`,
		},
		"import": {
			append:   true,
			pkgName:  "main",
			imports:  []string{"os"},
			typeName: "*os.File",
			optName:  "File",
			before: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

import "io"

// Reader
type OptionalReader struct {
	v   io.Reader
	has bool
}

func NoneReader() OptionalReader {
	return OptionalReader{}
}

func SomeReader(v io.Reader) OptionalReader {
	return OptionalReader{
		v:   v,
		has: true,
	}
}

func (o OptionalReader) Get() (Reader, bool) {
	return o.v, o.has
}
`,
			after: `// Code generated by github.com/acomagu/optional. DO NOT EDIT.

package main

import "os"

import "io"

// Reader
type OptionalReader struct {
	v   io.Reader
	has bool
}

func NoneReader() OptionalReader {
	return OptionalReader{}
}

func SomeReader(v io.Reader) OptionalReader {
	return OptionalReader{
		v:   v,
		has: true,
	}
}

func (o OptionalReader) Get() (Reader, bool) {
	return o.v, o.has
}

// File
type OptionalFile struct {
	v   *os.File
	has bool
}

func NoneFile() OptionalFile {
	return OptionalFile{}
}

func SomeFile(v *os.File) OptionalFile {
	return OptionalFile{
		v:   v,
		has: true,
	}
}

func (o OptionalFile) Get() (*os.File, bool) {
	return o.v, o.has
}
`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, err := ioutil.TempFile("", "opt.go")
			if err != nil {
				t.Fatal(errors.Wrap(err, "could not create temp file"))
			}

			if _, err := f.WriteString(test.before); err != nil {
				t.Fatal(errors.Wrap(err, "could not write to temp file"))
			}

			f.Seek(0, 0)

			g := &Generator{
				f:       f,
				append:  test.append,
				imports: test.imports,
				pkgName: test.pkgName,
				typName: test.typeName,
				optName: test.optName,
			}
			if err := g.generate(); err != nil {
				t.Fatal(err)
			}

			f.Seek(0, 0)
			res, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatal(errors.Wrap(err, "could not read the tmp file"))
			}

			if diff := cmp.Diff(strings.Split(string(res), "\n"), strings.Split(test.after, "\n")); diff != "" {
				t.Errorf("the result file has diff: %s", diff)
			}
		})
	}
}