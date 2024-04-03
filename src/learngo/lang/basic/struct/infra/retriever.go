package infra

import (
	"io"
	"net/http"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (Retriever) Retrieve(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(content)
}
