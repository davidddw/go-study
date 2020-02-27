// showSdkForGo project showSdkForGo.go
// examples  project examples.go
//	package main

//	import (
//		"fmt"
//		"showSdk"
//	)

//	func main() {
//		sa := showSdk.NewShowApi(162, "ade009*************e8d1c398")
//		sa.AddFile("src_img", "0.jpg")
//		sa.AddFile("logo_img", "0.jpg")
//		sa.Add("xy", "+50")
//		fmt.Println(sa.Post("http://route.showapi.com/1-2"))
//	}

package showSdk

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

func NewShowApi(appid int, sign string) *ShowApi {
	value := make(url.Values)
	value.Set("showapi_appid", strconv.Itoa(appid))
	return &ShowApi{appid, sign, value, make(map[string]string)}
}

//ShowApi结构体,Values储存普通表单数据,file储存文件数据
type ShowApi struct {
	appid int
	sign  string
	url.Values
	file map[string]string
}

//添加上传文件内容
func (sa *ShowApi) AddFile(formfile, filename string) {
	sa.file[formfile] = filename
}

//Set别名,拒绝key对多个value
func (sa *ShowApi) Add(key, value string) {
	sa.Set(key, value)
}

//以post方式发送一个调用请求,返回接口body,忽略file
func (sa *ShowApi) Get(url string) (string, error) {
	sa.sum()
	req := httplib.Get(url + "?" + sa.Encode())
	return req.String()
}

//以post方式发送一个调用请求,返回接口body,上传文件需用此方式
func (sa *ShowApi) Post(url string) (string, error) {
	sa.sum()
	req := httplib.Post(url)
	for k, v := range sa.Values {
		req.Param(k, v[0])
	}
	for k, v := range sa.file {
		req.PostFile(k, v)
	}
	return req.String()
}

//清空表单数据，用于新建请求
func (sa *ShowApi) Empty() {
	value := make(url.Values)
	value.Set("showapi_appid", sa.Values.Get("showapi_appid"))
	sa.Values = value
	sa.file = make(map[string]string)
}

//解析返回的json
func (sa *ShowApi) unmarshal(saRet map[string]interface{}) (map[string]interface{}, error) {
	if saRet == nil {
		return nil, errors.New("unmarshal error: showapi return body is nil")
	}
	if saRet["showapi_res_code"].(float64) != 0 {
		return nil, errors.New("showapi error: " + saRet["showapi_res_error"].(string))
	}
	return saRet["showapi_res_body"].(map[string]interface{}), nil

}

//签名加密
func (sa *ShowApi) sum() {
	prc, _ := time.LoadLocation("PRC")
	t := time.Now().In(prc)
	sa.Set("showapi_timestamp", fmt.Sprintf("%d%0.2d%0.2d%0.2d%0.2d%0.2d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
	keyArr := make([]string, len(sa.Values))
	for k, _ := range sa.Values {
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)
	valuestr := ""
	for _, k := range keyArr {
		for _, v := range sa.Values[k] {
			valuestr = valuestr + k + v
		}
	}
	valuestr += sa.sign
	sa.Set("showapi_sign", fmt.Sprintf("%X", md5.Sum([]byte(valuestr))))
}
