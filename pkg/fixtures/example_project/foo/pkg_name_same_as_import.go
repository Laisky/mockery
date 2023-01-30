package foo

import "github.com/Laisky/mockery/v2/pkg/fixtures/example_project/bar/foo"

type PackageNameSameAsImport interface {
	NewClient() foo.Client
}
