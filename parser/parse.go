package parser

import (
	"errors"
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

func formatHead() {
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
	fmt.Printf("|%s|%s|%s|\n", "ROOT"+strings.Repeat(" ", PARENT-len("ROOT")), "KEY"+strings.Repeat(" ", KEY-len("KEY")), "VALUE"+strings.Repeat(" ", VALUE-len("VALUE")))
}
func formatSideLine() {
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
}

func formatLine(Parent string, Key, Value string) {
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

// Print the struct to os.stdout
// TODO: support output type
func formatStructTable(t reflect.Type, v reflect.Value, Parent string, deepth int) {
	defer func() {
		if deepth == 0 {
			formatSideLine()
		}
	}()
	if deepth == 0 {
		formatHead()
	}

	formatSideLine()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Struct && deepth == 0 {
			formatStructTable(t.Field(i).Type, v.Field(i), t.Field(i).Name, 1)
			continue
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
				formatLine(Parent, key, value)
			} else {
				formatLine("", key, value)
			}
		}
	}
}

// Table show be a struct.
// if you want to use a struct pointer, you shold get the value first
// *T > T
func Parse(root interface{}) error {
	if reflect.TypeOf(root).Kind() != reflect.Struct {
		return errors.New("root is not struct")
	}
	formatStructTable(reflect.TypeOf(root), reflect.ValueOf(root), "", 0)
	return nil
}
