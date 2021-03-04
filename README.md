# qqinfo  
  qq 相关接口封装
------

## install
------
```
go get -u github.com/chinacase/qqinfo
```

## 获取QQ图像 昵称
------
```

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

func main() {
	r := qqinfo.CheckURL("www.uupf.com")
	fmt.Println(r)
}
```

