package qqinfo

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/axgle/mahonia"
	"github.com/tidwall/gjson"
)

//ResultInfo QQ信息结构体
type ResultInfo struct {
	Image string //qq图片
	Name  string //qq 昵称
}

//GetQQInfo 获取QQinfo
func GetQQInfo(qq string) (ressult ResultInfo, err error) {
	geturl := "http://r.qzone.qq.com/fcg-bin/cgi_get_portrait.fcg?get_nick=1&uins=" + qq
	resp, err := http.Get(geturl)
	ressult = ResultInfo{}
	if err != nil {
		err = errors.New("获取失败")
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	resHTML := mahonia.NewDecoder("gbk").ConvertString(string(body)) //
	jsonStr := resHTML[17 : len(resHTML)-1]

	res := gjson.Get(jsonStr, qq)

	if false == res.Exists() {
		err = errors.New("获取失败,请检查QQ号是否正确")
		return
	}
	ressult.Image = res.Array()[0].String()
	ressult.Name = res.Array()[6].String()
	return
}
