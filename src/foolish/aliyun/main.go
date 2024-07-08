package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

func main() {
	client, err := oss.New("endpoint", "AK", "AKS")
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.Bucket("BUCKET")
	if err != nil {
		log.Fatalln(err)
	}

	err = bucket.PutObjectFromFile("test-object", "C:/Users/Administrator/Desktop/20240704125216.png")
	if err != nil {
		log.Fatalln(err)
	}
}
