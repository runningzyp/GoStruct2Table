# GoStruct2Table

format your struct like a table.

```go
var p = struct {
	Name    string
	Age     int
	Address string
}{
	Name:    "xiangcai",
	Age:     18,
	Address: "shanghai",
}
```

```text
+----------+--------------------+---------------------------+
|ROOT      |KEY                 |VALUE                      |
+----------+--------------------+---------------------------+
|          |Name                |xiangcai                   |
|          |Age                 |18                         |
|          |Address             |shanghai                   |
+----------+--------------------+---------------------------+
```

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
	Name:    "xiangcai",
	Age:     18,
	Address: "shanghai",
	Nested:  Nested{Core: 1},
}
```

```text
+----------+--------------------+---------------------------+
|ROOT      |KEY                 |VALUE                      |
+----------+--------------------+---------------------------+
|          |Name                |xiangcai                   |
|          |Age                 |18                         |
|          |Address             |shanghai                   |
+----------+--------------------+---------------------------+
|Nested    |Core                |1                          |
+----------+--------------------+---------------------------+
```