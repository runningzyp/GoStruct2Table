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

type Row struct {
	key       string
	value     string
	parent    *Row
	childs    []*Row
	parentKey string // parent.childs[middle] == this
}

func formatHead() {
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
	fmt.Printf("|%s|%s|%s|\n", "ROOT"+strings.Repeat(" ", PARENT-len("ROOT")), "KEY"+strings.Repeat(" ", KEY-len("KEY")), "VALUE"+strings.Repeat(" ", VALUE-len("VALUE")))
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
}

func formatBottom() {
	fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
}

func formatLine() string {
	return fmt.Sprintf("+%s+%s+%s+\n", strings.Repeat("-", PARENT), strings.Repeat("-", KEY), strings.Repeat("-", VALUE))
}
func formatContent(Parent string, Key, Value string) string {
	if len(Parent) > PARENT {
		Parent = Parent[:PARENT]
	}
	if len(Key) > KEY {
		Key = Key[:KEY]
	}
	if len(Value) > VALUE {
		Value = Value[:VALUE]
	}
	return fmt.Sprintf("|%s|%s|%s|", Parent+strings.Repeat(" ", PARENT-len(Parent)), Key+strings.Repeat(" ", KEY-len(Key)), Value+strings.Repeat(" ", VALUE-len(Value)))
}

// Print the struct to os.stdout
// TODO: support output type
func parseStructTable(t reflect.Type, v reflect.Value, Parent *Row) {
	for i := 0; i < t.NumField(); i++ {
		var child = Row{parent: Parent}
		child.key = t.Field(i).Name
		Parent.childs = append(Parent.childs, &child)
		if t.Field(i).Type.Kind() == reflect.Struct && Parent.parent == nil {
			parseStructTable(t.Field(i).Type, v.Field(i), &child)
			continue
		}
		if !v.Field(i).CanInterface() {
			child.value = "<hidden field>"
			continue
		}
		switch v.Field(i).Interface().(type) {
		case string:
			child.value, _ = v.Field(i).Interface().(string)
		case int:
			ret, _ := v.Field(i).Interface().(int)
			child.value = strconv.Itoa(ret)
		case bool:
			ret, _ := v.Field(i).Interface().(bool)
			child.value = strconv.FormatBool(ret)
		default:
			child.value = fmt.Sprintf("%v", v.Field(i).Interface())
		}
		re := regexp.MustCompile(`\r|\n|\t|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`)
		child.value = re.ReplaceAllString(child.value, "")

	}
}

func formatStructTable(root *Row) {
	var outer string
	var inner string
	var line = formatLine()
	var lineFeed = "\n"
	for _, row := range root.childs {
		if row.childs == nil {
			inner += formatContent(row.parentKey, row.key, row.value) + lineFeed
		} else {
			if inner != "" {
				outer += inner + line
			}
			inner = ""
			row.childs[len(row.childs)/2].parentKey = row.key
			for _, child := range row.childs {
				inner += formatContent(child.parentKey, child.key, child.value) + lineFeed
			}
			outer += inner + line
			inner = ""
		}
	}
	formatHead()
	fmt.Printf(outer)

}

// Table show be a struct.
// if you want to use a struct pointer, you shold get the value first
// *T > T
func Parse(ins interface{}) error {
	if reflect.TypeOf(ins).Kind() != reflect.Struct {
		return errors.New("root is not struct")
	}
	var root = &Row{}
	parseStructTable(reflect.TypeOf(ins), reflect.ValueOf(ins), root)
	formatStructTable(root)
	return nil
}
