package examples

import (
	"strings"
	"testing"

	"github.com/Laisky/mockery/v2/examples/mocks"
	"github.com/Laisky/testify/require"
)

func Test_Hello(t *testing.T) {
	i := mocks.NewInterface(t)
	i.Catch("Hello", func(arguments ...interface{}) bool {
		return true
	}).Return(func(names ...string) string {
		return "hello: " + strings.Join(names, ",")
	})

	require.Equal(t, "hello: darth", i.Hello("darth"))
	require.Equal(t, "hello: darth,vadar", i.Hello("darth", "vadar"))
}
