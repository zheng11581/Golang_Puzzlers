package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirected to ", req.URL)
			return nil
		},
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	_, err = httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", s)
}
