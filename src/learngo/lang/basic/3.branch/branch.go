package main

import (
	"log"
	"os"
)

// if 条件语句不需要括号
// if 条件语句里面可以赋值变量，但是变量的作用域为block
func condition() {
	const filename string = "abc.txt"
	// if content, err := os.ReadFile(filename); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Fatalf("%s", content)
	// }
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s", content)
	}
}

// case语句自己break，如果不需要break需要fallthrough
// switch语句两种写法：
// 1. switch + 表达式
func evaluation(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		log.Fatalf("Wrong operation: %d", result)
	}
	return result
}

// 2. switch 没有表达式
func grade(score int) string {
	var grade string = ""
	switch {
	case score < 0 || score > 100:
		log.Fatalf("Wrong score: %d", score)
	case score < 60:
		grade = "不及格"
	case score < 80:
		grade = "良好"
	case score <= 100:
		grade = "优秀"
	}
	return grade
}

func main() {
	condition()
	log.Println(evaluation(3, 4, "+"))
	log.Println(grade(101))
}
