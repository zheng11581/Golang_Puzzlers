package main

import (
	"flag"
	"fmt"
)

func main() {
	// 类型推断的用途：提高代码的灵活性
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

func getTheFlag() *int {
	return flag.Int("number", 1, "The number of greeting object")
}

// 如果要改变flag的类型，只需要改变 getTheFlag 表达式的类型
// func getTheFlag() *string {
// 	return flag.String("name", "everybody", "The greeting object")
// }
