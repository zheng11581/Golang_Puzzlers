package main

import (
	"fmt"
	"time"
	"w2/struct/infra"
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
