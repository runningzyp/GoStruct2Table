# GoStruct2Table

format your struct like a table.

```go
var p = struct {
	Name    string
	Age     int
	Address string
}{
	Name:    "test",
	Age:     18,
	Address: "test",
}
```

| CONFIG | KEY             | VALUE |
| ----- | --------------- | ----- |
|        | Name            | test  |
| ROOT   | Age             | 18    |
|        | Address         | test  |
|        | <default: null> | KEYS: |

```go
type Nested struct {
	Core int
}
var p = struct {
	Name    string
	Age     int
	Address string
	Nested  Nested
}{
	Name:    "test",
	Age:     18,
	Address: "test",
	Nested:  Nested{Core: 1},
}
```
