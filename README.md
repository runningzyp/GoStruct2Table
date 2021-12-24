# GoStruct2Table

format your struct like a table.

## Installing

```
$ go get -u -v github.com/runningzyp/GoStruct2Table
```

## Simple Example



```go
import parser "github.com/runningzyp/GoStruct2Table"

var p = struct {
	Name    string
	Age     int
	Address string
}{
	Name:    "xiangcai",
	Age:     18,
	Address: "shanghai",
}

parser.Parse(p)

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
Nested:

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