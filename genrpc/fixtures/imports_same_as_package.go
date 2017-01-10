package test

import (
	"github.com/chris-skud/genrpc/genrpc/fixtures/test"
)

type C int

type ImportsSameAsPackage interface {
	A() test.B
	B() KeyManager
	C(C)
}
