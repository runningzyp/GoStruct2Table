package parser

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	TOTAL  = 100
	PARENT = 10
	KEY    = 20
	VALUE  = 70
)

func FormatHead() {
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
	fmt.Printf("|%s|%s|%s|\n", "ROOT"+strings.Repeat(" ", PARENT-len("ROOT")), "KEY"+strings.Repeat(" ", KEY-len("KEY")), "VALUE"+strings.Repeat(" ", VALUE-len("VALUE")))
}
func FormatSideLine() {
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
}

func FormatLine(Parent string, Key, Value string) {
	if len(Parent) > PARENT {
		Parent = Parent[:PARENT]
	}
	if len(Key) > KEY {
		Key = Key[:KEY]
	}
	if len(Value) > VALUE {
		Value = Value[:VALUE]
	}
	fmt.Printf("|%s|%s|%s|\n", Parent+strings.Repeat(" ", PARENT-len(Parent)), Key+strings.Repeat(" ", KEY-len(Key)), Value+strings.Repeat(" ", VALUE-len(Value)))
}

func FormatStructTable(t reflect.Type, v reflect.Value, Parent string, deepth int) {
	defer func() {
		if deepth == 0 {
			FormatSideLine()
		}
	}()
	if deepth == 0 {
		FormatHead()
	}

	FormatSideLine()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Struct && deepth == 0 {
			FormatStructTable(t.Field(i).Type, v.Field(i), t.Field(i).Name, 1)
			break
		}
		var key, value string
		key = t.Field(i).Name

		switch v.Field(i).Interface().(type) {
		case string:
			value, _ = v.Field(i).Interface().(string)
		case int:
			ret, _ := v.Field(i).Interface().(int)
			value = strconv.Itoa(ret)
		case bool:
			ret, _ := v.Field(i).Interface().(bool)
			value = strconv.FormatBool(ret)
		default:
			value = fmt.Sprintf("%v", v.Field(i).Interface())
		}
		re := regexp.MustCompile(`\r|\n|\t|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`)
		value = re.ReplaceAllString(value, "")
		if value != "" {
			if i == (t.NumField()-1)/2 {
				FormatLine(Parent, key, value)
			} else {
				FormatLine("", key, value)
			}
		}
	}
}
