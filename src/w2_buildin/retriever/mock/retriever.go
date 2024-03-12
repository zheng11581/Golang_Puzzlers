package mock

type Retriever struct {
	Content string
}

func (r *Retriever) Post(url string, form map[string]string) {
	r.Content = form["content"]
}

func (r *Retriever) Get(url string) string {
	return r.Content
}
