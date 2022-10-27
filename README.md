# tclient

基于 [req](https://github.com/imroc/req) 封装的灵活的腾讯云 SDK，没有结构化的 request 和 response，可直接参考 [官方 API 文档](https://cloud.tencent.com/document/api) 传入必要参数构造并发送请求。

## 示例

查询 [tcm](https://cloud.tencent.com/product/tcm) 所支持的地域列表，调用地域管理系统的云 API `DescribeRegions` 进行查询，参考 [这篇接口文档](https://cloud.tencent.com/document/api/1596/77930) 构造请求:

```go
package main

import (
	"github.com/imroc/tclient"
	"os"
)

func main() {
	secretId := os.Getenv("SECRET_ID")
	secretKey := os.Getenv("SECRET_KEY")
	if secretId == "" || secretKey == "" {
		panic("need secretId and secretKey")
	}
	client := tclient.NewClient("region", "2022-06-27").
		WithCredential(secretId, secretKey)
	client.DevMode()
	client.NewRequest(
		"DescribeRegions",
		map[string]interface{}{
			"Product": "tcm"}).
		Do()
}
```

运行一下看看请求和响应内容:

```bash
$ export SECRET_ID=***
$ export SECRET_KEY=***
$ go run .
2022/10/27 14:23:46.506785 DEBUG [req] HTTP/1.1 POST https://region.tencentcloudapi.com
POST / HTTP/1.1
Host: region.tencentcloudapi.com
User-Agent: req/v3 (https://github.com/imroc/req)
Content-Length: 17
Authorization: TC3-HMAC-SHA256 Credential=************************************/2022-10-27/region/tc3_request, SignedHeaders=content-type;host, Signature=****************************************************************
Content-Type: application/json; charset=utf-8
Nonce: 4726196768284630371
X-TC-Action: DescribeRegions
X-TC-Language: zh-CN
X-TC-RequestClient: SDK_GO_1.0.523
X-TC-Timestamp: 1666851825
X-TC-Version: 2022-06-27
Accept-Encoding: gzip

{"Product":"tcm"}
HTTP/1.1 200 OK
Server: nginx
Date: Thu, 27 Oct 2022 06:23:46 GMT
Content-Type: application/json
Content-Length: 3691
Connection: keep-alive

{"Response": {"TotalCount": 18, "RequestId": "10388d1a-44b3-45f6-b7c6-3da8e1373c9c", "RegionSet": [{"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-guangzhou", "RegionTypeMC": null, "RegionName": "\u534e\u5357\u5730\u533a(\u5e7f\u5dde)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-shenzhen-fsi", "RegionTypeMC": null, "RegionName": "\u534e\u5357\u5730\u533a(\u6df1\u5733\u91d1\u878d)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-shenzhen", "RegionTypeMC": null, "RegionName": "\u534e\u5357\u5730\u533a(\u6df1\u5733)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-shanghai", "RegionTypeMC": null, "RegionName": "\u534e\u4e1c\u5730\u533a(\u4e0a\u6d77)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-nanjing", "RegionTypeMC": null, "RegionName": "\u534e\u4e1c\u5730\u533a(\u5357\u4eac)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-hongkong", "RegionTypeMC": null, "RegionName": "\u6e2f\u6fb3\u53f0\u5730\u533a(\u4e2d\u56fd\u9999\u6e2f)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-beijing", "RegionTypeMC": null, "RegionName": "\u534e\u5317\u5730\u533a(\u5317\u4eac)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-tianjin", "RegionTypeMC": null, "RegionName": "\u534e\u5317\u5730\u533a(\u5929\u6d25)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-beijing-fsi", "RegionTypeMC": null, "RegionName": "\u534e\u5317\u5730\u533a(\u5317\u4eac\u91d1\u878d)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-singapore", "RegionTypeMC": null, "RegionName": "\u4e9a\u592a\u4e1c\u5357(\u65b0\u52a0\u5761)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "na-siliconvalley", "RegionTypeMC": null, "RegionName": "\u7f8e\u56fd\u897f\u90e8(\u7845\u8c37)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-chengdu", "RegionTypeMC": null, "RegionName": "\u897f\u5357\u5730\u533a(\u6210\u90fd)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-chongqing", "RegionTypeMC": null, "RegionName": "\u897f\u5357\u5730\u533a(\u91cd\u5e86)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "eu-frankfurt", "RegionTypeMC": null, "RegionName": "\u6b27\u6d32\u5730\u533a(\u6cd5\u5170\u514b\u798f)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "eu-moscow", "RegionTypeMC": null, "RegionName": "\u6b27\u6d32\u5730\u533a(\u83ab\u65af\u79d1)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-seoul", "RegionTypeMC": null, "RegionName": "\u4e9a\u592a\u4e1c\u5317(\u9996\u5c14)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "ap-tokyo", "RegionTypeMC": null, "RegionName": "\u4e9a\u592a\u4e1c\u5317(\u4e1c\u4eac)"}, {"RegionIdMC": null, "LocationMC": null, "RegionState": "AVAILABLE", "RegionNameMC": null, "Region": "na-ashburn", "RegionTypeMC": null, "RegionName": "\u7f8e\u56fd\u4e1c\u90e8(\u5f17\u5409\u5c3c\u4e9a)"}]}}
```

根据响应内容构造 struct 并 Unmarshal:

```go
package main

import (
	"fmt"
	"github.com/imroc/tclient"
	"os"
)

type DescribeRegionsResponse struct {
	Response struct {
		TotalCount int      `json:"TotalCount"`
		RequestId  string   `json:"RequestId"`
		RegionSet  []Region `json:"RegionSet"`
	} `json:"Response"`
}

type Region struct {
	RegionState string `json:"RegionState"`
	Region      string `json:"Region"`
	RegionName  string `json:"RegionName"`
}

func main() {
	secretId := os.Getenv("SECRET_ID")
	secretKey := os.Getenv("SECRET_KEY")
	if secretId == "" || secretKey == "" {
		panic("need secretId and secretKey")
	}
	client := tclient.NewClient("region", "2022-06-27").WithCredential(secretId, secretKey)
	var resp DescribeRegionsResponse
	product := "tcm"
	client.NewRequest(
		"DescribeRegions",
		map[string]interface{}{
			"Product": product}).
		SetResult(&resp).
		Do()
	fmt.Println(product + " supported regions:")
	for _, region := range resp.Response.RegionSet {
		fmt.Println(region.RegionName)
	}
}
```

运行一下：

```bash
$ export SECRET_ID=***
$ export SECRET_KEY=***
$ go run .
tcm supported regions:
华南地区(广州)
华南地区(深圳金融)
华南地区(深圳)
华东地区(上海)
华东地区(南京)
港澳台地区(中国香港)
华北地区(北京)
华北地区(天津)
华北地区(北京金融)
亚太东南(新加坡)
美国西部(硅谷)
西南地区(成都)
西南地区(重庆)
欧洲地区(法兰克福)
欧洲地区(莫斯科)
亚太东北(首尔)
亚太东北(东京)
美国东部(弗吉尼亚)

```