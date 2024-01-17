//package main
//
//import (
//	"fmt"
//	"github.com/oschwald/geoip2-golang"
//	"log"
//	"net"
//)
//
//func main() {
//	// 打开GeoIP2数据库文件
//	db, err := geoip2.Open("static/Location/GeoLite2-City.mmdb")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// 用户的IP地址
//	ip := net.ParseIP("2.16.28.239")
//
//	// 查询用户的地理位置
//	record, err := db.City(ip)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(record)
//
//	// 输出用户的地理位置信息
//	fmt.Printf("国家: %s\n", record.Country.Names["zh-CN"])
//	if len(record.Subdivisions) > 0 {
//		fmt.Printf("省份: %s\n", record.Subdivisions[0].Names["zh-CN"])
//	} else {
//		fmt.Println("省份信息不可用")
//	}
//	fmt.Printf("城市: %s\n", record.City.Names["zh-CN"])
//}

/*
调用百度地图api来获取ip地址的地理位置，需要安装第三方依赖：
go get github.com/menduo/gobaidumap
地址为：https://gitee.com/menduo/gobaidumap
*/
//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//)
//
//type Location struct {
//	Address string `json:"address"`
//	Content struct {
//		Address       string `json:"address"`
//		AddressDetail struct {
//			City     string `json:"city"`
//			CityCode int    `json:"city_code"`
//			Province string `json:"province"`
//		} `json:"address_detail"`
//		Point struct {
//			X string `json:"x"`
//			Y string `json:"y"`
//		} `json:"point"`
//	} `json:"content"`
//	Status int `json:"status"`
//}
//
//func main() {
//	// 此处填写您在控制台-应用管理-创建应用后获取的AK
//	//ak := os.Getenv("GOBAIDUMAP_AK")
//	ak := "v0YPDZ0mFKhUncims5v9lAWoVrCjigks"
//
//	// 服务地址
//	host := "https://api.map.baidu.com"
//
//	// 接口地址
//	uri := "/location/ip"
//
//	// 设置请求参数
//	params := url.Values{
//		"ip":   []string{"125.220.102.27"},
//		"coor": []string{"bd09ll"},
//		"ak":   []string{ak},
//	}
//
//	// 发起请求
//	request, err := url.Parse(host + uri + "?" + params.Encode())
//	if nil != err {
//		fmt.Printf("host error: %v", err)
//		return
//	}
//
//	resp, err := http.Get(request.String())
//	fmt.Printf("url: %s\n", request.String())
//	defer resp.Body.Close()
//	if err != nil {
//		fmt.Printf("request error: %v", err)
//		return
//	}
//	body, err2 := ioutil.ReadAll(resp.Body)
//	if err2 != nil {
//		fmt.Printf("response error: %v", err2)
//	}
//
//	// 解析 JSON 数据
//	var location Location
//	err = json.Unmarshal([]byte(body), &location)
//	if err != nil {
//		fmt.Println("解析 JSON 数据失败：", err)
//		return
//	}
//
//	// 输出详细地址信息
//	fmt.Println("详细地址信息：", location.Address)
//
//	// 输出省份、城市
//	fmt.Println("省份：", location.Content.AddressDetail.Province)
//	fmt.Println("城市：", location.Content.AddressDetail.City)
//
//	//fmt.Println(string(body))
//
//	return
//}

/*
腾讯地图ip定位
*/

package main

import (
	"encoding/json"
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

func main() {
	core.InitConf()

	// 此处填写您在控制台-应用管理-创建应用后获取的AK
	key := global.Config.TxMap.Key

	// 服务地址
	host := "https://apis.map.qq.com/ws/location"

	// 接口地址
	uri := "/v1/ip"

	var ip string
	fmt.Printf("请输入你的ip：")
	fmt.Scanf("%s", &ip)

	// 设置请求参数
	params := url.Values{
		"ip":  []string{ip},
		"key": []string{key},
	}

	// 发起请求
	request, err := url.Parse(host + uri + "?" + params.Encode())
	if nil != err {
		fmt.Printf("host error: %v", err)
		return
	}

	resp, err := http.Get(request.String())
	//fmt.Printf("url: %s\n", request.String())
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("request error: %v", err)
		return
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Printf("response error: %v", err2)
	}

	// 解析 JSON 数据
	var location Location
	err = json.Unmarshal([]byte(body), &location)
	if err != nil {
		fmt.Println("解析 JSON 数据失败：", err)
		return
	}

	addr := "局域网IP"

	// 输出省份、城市
	if location.Result.AdInfo.Province != "" {
		addr = ""
		addr += location.Result.AdInfo.Province
	}
	if location.Result.AdInfo.City != "" {
		addr += "，" + location.Result.AdInfo.City
	}
	if location.Result.AdInfo.District != "" {
		addr += "，" + location.Result.AdInfo.District
	}

	fmt.Printf("您的IP所在位置为：%s", addr)
}
