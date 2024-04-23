package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

func main() {
	/**
	 * 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	 * 此处以把AccessKey和AccessKeySecret保存在环境变量为例说明。您也可以根据业务需要，保存到配置文件里。
	 * 强烈建议不要把AccessKey和AccessKeySecret保存到代码里，会存在密钥泄漏风险
	 */
	//AccessKeyId := os.Getenv("NLP_AK_ENV")
	//AccessKeySecret := os.Getenv("NLP_SK_ENV")
	//client, err := sdk.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessKeySecret)
	//if err != nil {
	//	panic(err)
	//}
	//request := requests.NewCommonRequest()
	//request.Domain = "alinlp.cn-hangzhou.aliyuncs.com"
	//request.Version = "2020-06-29"
	//// 因为是RPC接口，因此需指定ApiName(Action)
	//request.ApiName = "GetWsChGeneral"
	//request.QueryParams["ServiceCode"] = "alinlp"
	//request.QueryParams["Text"] = "这是一双很棒的运动鞋"
	//request.QueryParams["TokenizerId"] = "GENERAL_CHN"
	//response, err := client.ProcessCommonRequest(request)
	//if err != nil {
	//	panic(err)
	//}
	//txt := response.GetHttpContentString()
	txt := `{"RequestId":"894BDD97-26CB-5945-B3EF-7CF45B2CFA65","Data":{"result":[{"id":"0","word":"这","tags":["基本词-中文"]},{"id":"1","word":"是","tags":["基本词-中文","文体娱乐类-flash作品"]},{"id":"2","word":"一","tags":["基本词-中文","文体娱乐类-游戏类","文体娱乐类-戏剧歌曲类","文体娱乐类-书文课程类","文体娱乐类-影视类","体育-独立无后缀球队"]},{"id":"3","word":"双","tags":["基本词-中文","产品类型修饰词"]},{"id":"4","word":"很","tags":["基本词-中文"]},{"id":"5","word":"棒","tags":["产品类型-简单","基本词-中文","特殊商品-收藏"]},{"id":"6","word":"的","tags":["基本词-中文","文体娱乐类-flash作品"]},{"id":"7","word":"运动鞋","tags":["产品类型-简单","基本词-中文","产品类型修饰词"]}],"success":true}}`
	parseNLP([]byte(txt))
}

func marshal() {
	o := Order{
		ID:         "1",
		Name:       "John Smith",
		Quantity:   100,
		TotalPrice: 99.99,
	}
	//fmt.Printf("%+vn", o)
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+vn", string(b))
}

func unmarshal() {
	var o Order
	err := json.Unmarshal([]byte(`{"id":"1","name":"John Smith","quantity":100,"total_price":99.99}`), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

type WsChResponse struct {
	Data struct {
		Result []struct {
			ID   string   `json:"id"`
			Word string   `json:"word"`
			Tags []string `json:"tags"`
		} `json:"result"`
		Success bool `json:"success"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type WsChRequest struct {
	Action      string `json:"action"`
	Text        string `json:"text"`
	ServiceCode string `json:"service_code"`
	TokenizerId string `json:"tokenizer_id"`
	OutType     string `json:"out_type"`
}

func parseNLP(res []byte) WsChResponse {

	var resp WsChResponse

	err := json.Unmarshal(res, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
	return resp

}
