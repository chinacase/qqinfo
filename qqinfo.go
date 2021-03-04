package qqinfo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	"github.com/tidwall/gjson"
)

//ResultInfo QQ信息结构体
type ResultInfo struct {
	Image string //qq图片
	Name  string //qq 昵称
}

const (
	//StatusUnknown 安全性未知
	StatusUnknown = 1
	//StatusPass 安全网站
	StatusPass = 3
	//StatusNoPass 危险网站
	StatusNoPass = 2
)

//Result Result 解析
type Result struct {
	Status int    ///状态 1：安全性未知 2：危险网站 3：安全网站
	Msg    string //提示
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

//CheckURL qq域名检测
func CheckURL(urlString string) (result Result) {
	var getURL string
	result = Result{}

	if !strings.HasPrefix(urlString, "http://") && !strings.HasPrefix(urlString, "https://") {
		urlString = "http://" + urlString
	}
	getURL = "https://cgi.urlsec.qq.com/index.php?m=check&a=check&url=" + urlString
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, _ := http.NewRequest("GET", getURL, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4406.0 Safari/537.36")
	req.Header.Set("Referer", "https://guanjia.qq.com")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	resHTML := string(body)
	jsonStr := resHTML[1 : len(resHTML)-1]

	res := gjson.Get(jsonStr, "data.results")

	if false == res.Exists() {
		result.Status = StatusUnknown
		result.Msg = "未知"
		return
	}
	r := res.Map()
	if r["whitetype"].Int() == 3 || r["whitetype"].Int() == 1 {
		result.Status = StatusPass
		result.Msg = "网站安全"
		return
	}

	result.Status = StatusNoPass
	result.Msg = r["WordingTitle"].String()
	return
}
