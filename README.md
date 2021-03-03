# qqinfo  
  获取QQ图像 昵称
------
# install
------
```
go get -u github.com/chinacase/qqinfo
```
```package main

import (
	"fmt"

	"github.com/chinacase/qqinfo"
)
# example
------
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
