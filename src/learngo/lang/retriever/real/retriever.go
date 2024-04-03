package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Post(url string, form map[string]string) {
	r.UserAgent = form["userAgent"]
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(body)
}
