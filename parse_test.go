package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
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
	Parse(test)

}

func TestParseNested(t *testing.T) {
	type Nested struct {
		Height int
		Weight int
	}
	var n = struct {
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
	Parse(n)
}

func TestParseNestedAnym(t *testing.T) {
	type test struct {
		Height int
		Weight int
	}

	var n = struct {
		Name    string
		Age     int
		Address string
		test
	}{
		Name:    "test",
		Age:     18,
		Address: "test",
		test:    test{Height: 170, Weight: 60},
	}
	Parse(n)
}
