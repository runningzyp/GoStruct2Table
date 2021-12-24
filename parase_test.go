package parser

import (
	"testing"
)

type Nested struct {
	Core int
}

func TestParase(t *testing.T) {
	var test = struct {
		Name    string
		Age     int
		Address string
	}{
		Name:    "test",
		Age:     18,
		Address: "test",
	}
	Parase(test)

}
func TestParaseNested(t *testing.T) {

	var n = struct {
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
	Parase(n)
}
