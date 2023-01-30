## Install

```sh
go install github.com/Laisky/mockery/v2@9b58e70
mockery --with-expecter --name="Interface"
```

## New Feature

### `.Catch` Mock method with any arguments

```go
func Test_Hello(t *testing.T) {
    i := mocks.NewInterface(t)
    // Catch can catch any method by name no matter what the arguments is.
    i.Catch("Hello", func(arguments ...interface{}) bool {
        return true
    }).Return(func(names ...string) string {
        return "hello: " + strings.Join(names, ",")
    })

    require.Equal(t, "hello: darth", i.Hello("darth"))
    require.Equal(t, "hello: darth,vadar", i.Hello("darth", "vadar"))
}
```
