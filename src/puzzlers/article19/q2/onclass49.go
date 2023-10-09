package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter funtion main.")
	caller1()
	fmt.Println("Exist funtion main.")
}

func caller1() {
	fmt.Println("Enter funtion caller1.")
	panic(fmt.Errorf("something error")) // panic() 函数的参数最好是一个error类型
	panic("something error")             // 不建议
	fmt.Println("Exist funtion caller1.")
}
