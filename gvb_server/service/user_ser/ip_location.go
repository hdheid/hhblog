package user_ser

import (
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Location 响应数据
type Location struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		IP       string `json:"ip"`
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		AdInfo struct {
			Nation   string `json:"nation"`
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
			Adcode   int    `json:"adcode"`
		} `json:"ad_info"`
	} `json:"result"`
}

// GetLocation 这个是调用腾讯的api，每天只能调用一万次
func GetLocation(ip string) (loc string) {
	key := global.Config.TxMap.Key

	// 服务地址
	host := "https://apis.map.qq.com/ws/location"

	// 接口地址
	uri := "/v1/ip"

	// 设置请求参数
	params := url.Values{
		"ip":  []string{ip},
		"key": []string{key},
	}

	// 发起请求
	request, err := url.Parse(host + uri + "?" + params.Encode())
	if nil != err {
		global.Log.Error("host error: %v", err)
		return
	}

	resp, err := http.Get(request.String()) //发起请求
	fmt.Printf("url: %s\n", request.String())
	defer resp.Body.Close()
	if err != nil {
		global.Log.Error("request error: %v", err)
		return
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		global.Log.Error("response error: %v", err2)
	}

	// 解析 JSON 数据
	var location Location
	err = json.Unmarshal([]byte(body), &location)
	if err != nil {
		global.Log.Error("解析 JSON 数据失败：", err)
		return
	}

	// 输出省份、城市
	if location.Result.AdInfo.Province != "" {
		fmt.Println("省份：", location.Result.AdInfo.Province)
		loc += location.Result.AdInfo.Province
	}
	if location.Result.AdInfo.City != "" {
		fmt.Println("城市：", location.Result.AdInfo.City)
		loc += "，" + location.Result.AdInfo.City
	}
	if location.Result.AdInfo.District != "" {
		fmt.Println("区：：", location.Result.AdInfo.District)
		loc += "，" + location.Result.AdInfo.District
	}

	return
}
