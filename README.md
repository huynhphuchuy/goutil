# goutilities
Go utilities

Example:

```go
package main

import (
	"PROJECTS/goutilities/Util"
	"fmt"
)

func main() {
	stringUtil := Util.InitString("huynhphuchuy")
	fmt.Println(stringUtil.Reverse())
	stringBetwwens, _ := stringUtil.Between("huynhphuchuy", "huynh", "huy")
	fmt.Println(stringBetwwens[0])
	stringConcat, _ := stringUtil.Concat("handsome", 69, 66.99, []string{"x", "x", "x"}, []int{6, 9, 9, 6})
	fmt.Println(stringConcat)
}
```
