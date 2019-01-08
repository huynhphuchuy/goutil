package Util

import (
	"goutilities/StringUtil"
)

type Utility interface {
	Reverse() string
	Between(string, string, string) ([]string, error)
	Concat(...interface{}) (string, error)
}

func InitString(string string) Utility {
	return Utility(StringUtil.String{Value: string})
}
