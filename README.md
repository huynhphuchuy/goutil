# goutilities
Go utilities

Example:

```go
package main

import (
	"fmt"

	"github.com/huynhphuchuy/goutilities/Util"
)

func main() {

	// Init String Util
	stringUtil := Util.InitString("huynhphuchuy")

	// Reverse string
	fmt.Println(stringUtil.Reverse())

	// Get string between
	stringBetweens, _ := stringUtil.Between("huynh", "huy")
	fmt.Println(stringBetweens[0])

	// Concat string with string, int, float, array
	stringConcat, _ := stringUtil.Concat("handsome", 69, 66.99, []string{"x", "x", "x"}, []int{6, 9, 9, 6})
	fmt.Println(stringConcat)

	// Check string is palindrome
	palindrome := Util.InitString("huynhuhuhnyuh")
	fmt.Println(stringUtil.IsPalindrome())
	fmt.Println(palindrome.IsPalindrome())

	// Randomize string
	fmt.Println(stringUtil.Random(69))

	// Remove String Diacritics
	diacriticsString := Util.InitString("Xin chào thế giới")
	fmt.Println(diacriticsString.RemoveDiacritics())
}
```
