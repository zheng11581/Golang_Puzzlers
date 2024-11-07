package main

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"os"
)

func main() {
	//推荐通过环境变量获取AKSK，这里也可以使用其他外部引入方式传入，如果使用硬编码可能会存在泄露风险。
	//您可以登录访问管理控制台获取访问密钥AK/SK，获取方式请参见https://support.huaweicloud.com/usermanual-ca/ca_01_0003.html。
	ak := os.Getenv("AccessKeyID")
	sk := os.Getenv("SecretAccessKey")
	// 【可选】如果使用临时AK/SK和SecurityToken访问OBS，同样建议您尽量避免使用硬编码，以降低信息泄露风险。您可以通过环境变量获取访问密钥AK/SK，也可以使用其他外部引入方式传入。
	securityToken := os.Getenv("SecurityToken")
	// endpoint填写Bucket对应的Endpoint, 这里以华北-北京四为例，其他地区请按实际情况填写。
	endPoint := "https://obs.cn-north-4.myhuaweicloud.com"
	// 创建obsClient实例
	// 如果使用临时AKSK和SecurityToken访问OBS，需要在创建实例时通过obs.WithSecurityToken方法指定securityToken值。
	obsClient, err := obs.New(ak, sk, endPoint, obs.WithSecurityToken(securityToken))
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
	}
	input := &obs.PutFileInput{}
	// 指定存储桶名称
	input.Bucket = "examplebucket"
	// 指定上传对象，此处以 example/objectname 为例。
	input.Key = "example/objectname"
	// 指定本地文件，此处以localfile为例
	input.SourceFile = "localfile"
	// 文件上传
	output, err := obsClient.PutFile(input)
	if err == nil {
		fmt.Printf("Put file(%s) under the bucket(%s) successful!\n", input.Key, input.Bucket)
		fmt.Printf("StorageClass:%s, ETag:%s\n",
			output.StorageClass, output.ETag)
		return
	}
	fmt.Printf("Put file(%s) under the bucket(%s) fail!\n", input.Key, input.Bucket)
	if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println("An ObsError was found, which means your request sent to OBS was rejected with an error response.")
		fmt.Println(obsError.Error())
	} else {
		fmt.Println("An Exception was found, which means the client encountered an internal problem when attempting to communicate with OBS, for example, the client was unable to access the network.")
		fmt.Println(err)
	}
}
