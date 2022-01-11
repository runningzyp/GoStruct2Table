package main

import (
	"github.com/runningzyp/GoStruct2Table/parser"
)

func main() {
	type Nested struct {
		Height int
		Weight int
	}
	var test = struct {
		Name    string
		Age     int
		Address string
		Nested  Nested
	}{
		Name:    "test",
		Age:     18,
		Address: "test",
		Nested:  Nested{Height: 170, Weight: 60},
	}
	parser.Parse(test)
}
