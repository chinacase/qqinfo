# qqinfo  
  qq 相关接口封装
------
# 获取QQ图像 昵称
## install
------
```
go get -u github.com/chinacase/qqinfo
```

## example
------
```
package main

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
------
## 域名报红检测
```
package main

import (
	"fmt"

	"github.com/chinacase/qqinfo"
)
func main() {
	r := qqinfo.CheckURL("www.uupf.com")
	fmt.Println(r)
}
```

