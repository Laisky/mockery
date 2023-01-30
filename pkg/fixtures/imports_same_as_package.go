package test

import (
	"github.com/Laisky/mockery/v2/pkg/fixtures/test"
)

type C int

type ImportsSameAsPackage interface {
	A() test.B
	B() KeyManager
	C(C)
}
