package table

import (
	"errors"
	"reflect"
	"table/parser"
)

// Table show be a struct.
// if you want to use a struct pointer, you shold get the value first
// *T > T
func Parse(root interface{}) error {
	if reflect.TypeOf(root).Kind() != reflect.Struct {
		return errors.New("root is not struct")
	}
	parser.FormatStructTable(reflect.TypeOf(root), reflect.ValueOf(root), "", 0)
	return nil
}
