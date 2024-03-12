package main

import (
	"fmt"
	"time"
	"w2/retriever/mock"
	"w2/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.immoc.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Content: ", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mock.Retriever{
		Content: "this is a fake immoc.com",
	}
	//fmt.Println(download(r))
	//fmt.Printf("%T %v\n", r, r)
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Content)
	} else {
		fmt.Println("not a mock retriever")
	}

	r = &real.Retriever{
		UserAgent: "Mozilla: 5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(download(r))
	//fmt.Printf("%T %v\n", r, r)
	inspect(r)

	// Type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a real retriever")
	}
}
