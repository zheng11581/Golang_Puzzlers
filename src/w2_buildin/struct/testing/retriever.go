package testing

type Retriever struct {
}

func (Retriever) Retrieve(url string) string {
	return "The Fake Content"
}
