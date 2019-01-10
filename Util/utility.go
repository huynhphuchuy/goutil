package Util

import (
	"github.com/huynhphuchuy/goutilities/StringUtil"
)

type Utility interface {
	Reverse() string
	Between(string, string) ([]string, error)
	Concat(...interface{}) (string, error)
	IsPalindrome() bool
	Random(int) string
}

func InitString(string string) Utility {
	return Utility(StringUtil.String{Value: string})
}
