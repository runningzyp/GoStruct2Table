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

var Lock = struct {
	title int
}{
	title: 0,
}

func FormatHead() {
	if Lock.title != 0 {
		return
	}
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
	fmt.Printf("|%s|%s|%s|\n", "CONFIG"+strings.Repeat(" ", PARENT-len("CONFIG")), "KEY"+strings.Repeat(" ", KEY-len("KEY")), "VALUE"+strings.Repeat(" ", VALUE-len("VALUE")))
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
	Lock.title = 1
}
func FormatLine(Parent string, isTail bool, Key, Value string) {
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
	if isTail {
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
	} else {
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat(" ", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
	}
}

func FormatStructTable(t reflect.Type, v reflect.Value, Parent string, deepth int) {
	defer func() {
		if deepth == 0 {
			FormatLine("", true, "<default: null>", "KEYS: ")
		}
	}()
	FormatHead()
	useDefault := []string{}
	var par string
	var hasPar bool
	isTail := false
	for i := 0; i < t.NumField(); i++ {

		if t.Field(i).Type.Kind() == reflect.Struct && deepth < 1 {
			FormatStructTable(t.Field(i).Type, v.Field(i), t.Field(i).Name, 1)
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
		if i == t.NumField()/2 {
			par = Parent
		}
		if value == "" {
			useDefault = append(useDefault, key)
		}
		if value != "" {
			FormatLine(par, isTail, key, value)
			if par != "" {
				hasPar = true
				par = ""
			}
		}
	}
	if !hasPar {
		par = Parent
	}
	if deepth != 0 {
		FormatLine(par, true, "<default: null>", "KEYS: "+strings.Join(useDefault, ","))
	}
}
