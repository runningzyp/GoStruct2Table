package parser

import "reflect"

func Parse(root interface{}) {
	FormatStructTable(reflect.TypeOf(root), reflect.ValueOf(root), "", 0)
}
