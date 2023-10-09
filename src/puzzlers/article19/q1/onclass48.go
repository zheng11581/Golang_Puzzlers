package main

import "fmt"

// panic 按照调用链反方向抛出运行时恐慌（打印运行时错误）

func main() {
	fmt.Println("Enter funtion main.")
	caller1()
	fmt.Println("Exist funtion main.")
}

func caller1() {
	fmt.Println("Enter funtion caller1.")
	caller2()
	fmt.Println("Exist funtion caller1.")
}

func caller2() {
	fmt.Println("Enter funtion caller2.")
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
	fmt.Println("Exist funtion caller2.")
}
