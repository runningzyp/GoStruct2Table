package parser

import "reflect"

func Parase(root interface{}) {
	FormatStructTable(reflect.TypeOf(root), reflect.ValueOf(root), "", 0)
}
