package main

import "fmt"

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)

}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head

}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func main() {
	var queue Queue
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())

	queue.Push("abc")
	fmt.Println(queue.Pop())
}
