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
	if err := Parse(test); err != nil {
		t.Error(err)
	}

}

func TestParseNested(t *testing.T) {
	type Nested struct {
		Height int
		Weight int
	}
	type Nested1 struct {
		Height int
		Weight int
	}
	var n = struct {
		Nested1 Nested1
		Name    string
		Age     int
		Address string
		Nested  Nested
		b       bool
	}{
		Name:    "test",
		Age:     18,
		Address: "test",
		Nested:  Nested{Height: 170, Weight: 60},
		Nested1: Nested1{Height: 170, Weight: 60},
		b:       true,
	}
	if err := Parse(n); err != nil {
		t.Error(err)
	}

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

func TestParseNotStruct(t *testing.T) {
	var e = map[string]string{"1": "1", "2": "2"}
	if err := Parse(e); err == nil {
		t.Error()
	}
}
