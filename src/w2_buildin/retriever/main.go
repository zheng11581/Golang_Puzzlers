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

type Poster interface {
	Post(url string, form map[string]string)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.immoc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(p Poster) {
	p.Post(url, map[string]string{
		"content":   "this is a another faked immoc.com",
		"userAgent": "IE 7.0",
	})
}

func session(rp RetrieverPoster) {
	post(rp)
	download(rp)
}

func inspect(r Retriever) {
	fmt.Printf("Inspecting: %T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("> Content: ", v.Content)
	case *real.Retriever:
		fmt.Println("> UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r RetrieverPoster
	r = &mock.Retriever{
		Content: "this is a fake immoc.com",
	}
	//fmt.Println(download(r))
	//fmt.Printf("%T %v\n", r, r)
	inspect(r)
	session(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Printf("> Content: %s\n", mockRetriever.Content)
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
	session(r)

	// Type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Printf("> UserAgent: %s\n", realRetriever.UserAgent)
	} else {
		fmt.Println("not a real retriever")
	}
}
