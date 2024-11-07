package main

// 引入依赖包
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
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
	obsClient, err := obs.New(ak, sk, endPoint, obs.WithSignature(obs.SignatureObs) /*, obs.WithSecurityToken(securityToken)*/)
	if err == nil {
		// 使用访问OBS

		// 关闭obsClient
		obsClient.Close()
	}
}
