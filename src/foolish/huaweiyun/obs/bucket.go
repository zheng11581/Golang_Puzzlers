package main

import (
	"fmt"
	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"os"
)

func main() {
	//推荐通过环境变量获取AKSK，这里也可以使用其他外部引入方式传入，如果使用硬编码可能会存在泄露风险。
	//您可以登录访问管理控制台获取访问密钥AK/SK，获取方式请参见https://support.huaweicloud.com/usermanual-ca/ca_01_0003.html。
	ak := os.Getenv("AccessKeyID")
	sk := os.Getenv("SecretAccessKey")
	// 【可选】如果使用临时AK/SK和SecurityToken访问OBS，同样建议您尽量避免使用硬编码，以降低信息泄露风险。您可以通过环境变量获取访问密钥AK/SK，也可以使用其他外部引入方式传入。
	// securityToken := os.Getenv("SecurityToken")
	// endpoint填写Bucket对应的Endpoint, 这里以华北-北京四为例，其他地区请按实际情况填写。
	endPoint := "https://obs.cn-north-4.myhuaweicloud.com"
	// 创建obsClient实例
	// 如果使用临时AKSK和SecurityToken访问OBS，需要在创建实例时通过obs.WithSecurityToken方法指定securityToken值。
	obsClient, err := obs.New(ak, sk, endPoint /*, obs.WithSecurityToken(securityToken)*/)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
	}
	input := &obs.CreateBucketInput{}
	// 指定存储桶名称
	input.Bucket = "examplebucket"
	// 指定存储桶所在区域，此处以“cn-north-4”为例，必须跟传入的Endpoint中Region保持一致。
	input.Location = "cn-north-4"
	// 指定存储桶的权限控制策略，此处以obs.AclPrivate为例。
	input.ACL = obs.AclPrivate
	// 指定存储桶的存储类型，此处以obs.StorageClassWarm为例。如果未指定该参数，则创建的桶为标准存储类型。
	input.StorageClass = obs.StorageClassWarm
	// 指定存储桶的AZ类型，此处以“3AZ”为例。不携带时默认为单AZ，如果对应region不支持多AZ存储，则该桶的存储类型仍为单AZ。
	input.AvailableZone = "3az"
	// 创建桶
	output, err := obsClient.CreateBucket(input)
	if err == nil {
		fmt.Printf("Create bucket:%s successful!\n", input.Bucket)
		fmt.Printf("RequestId:%s\n", output.RequestId)
		return
	}
	fmt.Printf("Create bucket:%s fail!\n", input.Bucket)
	if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println("An ObsError was found, which means your request sent to OBS was rejected with an error response.")
		fmt.Println(obsError.Error())
	} else {
		fmt.Println("An Exception was found, which means the client encountered an internal problem when attempting to communicate with OBS, for example, the client was unable to access the network.")
		fmt.Println(err)
	}
}
