package parser

import (
	"errors"
	"reflect"
)

func Parse(root interface{}) error {
	if reflect.TypeOf(root).Kind() != reflect.Struct {
		return errors.New("root is not struct")
	}
	FormatStructTable(reflect.TypeOf(root), reflect.ValueOf(root), "", 0)
	return nil
}
