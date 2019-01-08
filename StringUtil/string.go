package StringUtil

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type String struct {
	Value string
}

func (input String) Reverse() (output string) {
	for _, v := range input.Value {
		defer func(v rune) { output += string(v) }(v)
	}
	return output
}

func (input String) Between(raw string, start string, end string) ([]string, error) {
	re := regexp.MustCompile(fmt.Sprintf("%s%s%s", regexp.QuoteMeta(start), "(.*)", regexp.QuoteMeta(end)))
	if matches := re.FindStringSubmatch(raw); len(matches) > 1 {
		return matches[1:], nil
	}
	return []string{}, errors.New("No substring found!")
}

func (input String) Concat(params ...interface{}) (output string, error error) {
	output = input.Value
	for _, v := range params {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			slice := reflect.ValueOf(v)
			for i := 0; i < slice.Len(); i++ {
				output += fmt.Sprintf("%v", slice.Index(i))
			}
			break
		default:
			output += fmt.Sprintf("%v", v)

		}
	}
	return output, nil
}
