package main

import (
	"fmt"
	"sync"

	"github.com/runningzyp/GoStruct2Table/parser"
)

func test1() {
	fmt.Print("1")
}
func test2() {
	fmt.Print("2")
}

var a = `        +----------+--------------------+---------------------------+
        |ROOT      |KEY                 |VALUE                      |
        +----------+--------------------+---------------------------+
        |          |Name                |xiangcai                   |
        |          |Age                 |18                         |
        |          |Address             |shanghai                   |
        +----------+--------------------+---------------------------+
        +----------+--------------------+---------------------------+
        |Nested    |Core                |1                          |
        +----------+--------------------+---------------------------+`

func One(once sync.Once) {
	fmt.Print(a)
	once.Do(test1)
	once.Do(test1)
	once.Do(test2)
}

func Test() {
	var a = 1
	print(a)
	Test()
}

func main() {
	var test = struct {
		Name    string
		Age     int
		Address string
		Country string
	}{
		Name:    "test",
		Age:     18,
		Address: "shanghai",
		Country: "china",
	}
	parser.Parse(test)
}
