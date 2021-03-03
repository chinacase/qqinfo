# qqinfo
获取QQ图像 昵称
------
```package main

import (
	"fmt"

	"github.com/chinacase/qqinfo"
)

func main() {
	qq := "397932843"
	q, err := qqinfo.GetQQInfo(qq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q.Name)
	fmt.Println(q.Image)
}

```
