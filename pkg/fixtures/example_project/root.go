package example_project

import "github.com/Laisky/mockery/v2/pkg/fixtures/example_project/foo"

type Root interface {
	TakesBaz(*foo.Baz)
	ReturnsFoo() (foo.Foo, error)
}
