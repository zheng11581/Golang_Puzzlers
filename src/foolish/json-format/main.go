package main

import (
	"encoding/json"
	"fmt"
)

type OpsvAppInfo struct {
	AppId          string   `json:"app_id"`
	AppName        string   `json:"pipeline_name"`
	PipelineCode   string   `json:"pipeline_code"`
	Dc             string   `json:"dc"`
	Env            string   `json:"env"`
	Cmd            string   `json:"cmd"`
	DiversionAppId []string `json:"diversion_app_id"`
}

func main() {
	opsv := OpsvAppInfo{
		AppId:          "appid",
		AppName:        "appname",
		PipelineCode:   "pipeline_code",
		Dc:             "dc",
		Env:            "env",
		Cmd:            "cmd",
		DiversionAppId: []string{"diversion_app_id"},
	}
	data, _ := json.Marshal(opsv)
	fmt.Printf("JSON: %s\n", string(data))
	fmt.Printf("Struct: %+v\n", opsv) // 结论是将结构体转换为json后原结构体不变

}
