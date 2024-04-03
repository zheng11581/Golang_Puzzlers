package main

import (
	"fmt"
	"time"
	"zheng11581/learngo/lang/container/struct/infra"
)

func getRetriever() Retriever {
	return infra.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   5 * time.Minute,
	}
}

type Retriever interface {
	Retrieve(url string) string
}

func main() {
	var retriever Retriever
	retriever = getRetriever()
	fmt.Println(retriever.Retrieve("https://www.immoc.com"))
}
