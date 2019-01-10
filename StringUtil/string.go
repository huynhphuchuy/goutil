package StringUtil

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type String struct {
	Value string
}

// ******** INternal Functions *********

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// *************************************

func (input String) Reverse() (output string) {
	for _, v := range input.Value {
		defer func(v rune) { output += string(v) }(v)
	}
	return output
}

func (input String) Between(start string, end string) ([]string, error) {
	re := regexp.MustCompile(fmt.Sprintf("%s%s%s", regexp.QuoteMeta(start), "(.*)", regexp.QuoteMeta(end)))
	if matches := re.FindStringSubmatch(input.Value); len(matches) > 1 {
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

func (input String) IsPalindrome() bool {
	return input.Reverse() == input.Value
}

func (input String) Random(length int) string {
	runes := []rune(input.Value)
	randString := make([]rune, length)
	for i := range randString {
		randString[i] = runes[rand.Intn(len(runes))]
	}
	return string(randString)
}

func (input String) RemoveDiacritics() string {
	output := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(output, input.Value)
	return result
}
