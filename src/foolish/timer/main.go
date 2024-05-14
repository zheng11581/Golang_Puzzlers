package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	students := `[{"name":"John","age":30},{"name":"Jane","age":20}]`
	var result []Student
	json.Unmarshal([]byte(students), &result)
	ch := Consumer("Class 1")
	for _, student := range result {
		ch <- &student
	}
	<-time.After(5 * time.Second)

}

type Student struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Age   int    `json:"age"`
}

func (s Student) String() string {
	return fmt.Sprintf("[name: %s, class: %s, age: %d)]", s.Name, s.Class, s.Age)
}

func Consumer(name string) chan *Student {
	ch := make(chan *Student)
	go func() {
		for student := range ch {
			fmt.Printf("Consumer %s: %s\n", name, student.String())
		}
	}()
	return ch
}
