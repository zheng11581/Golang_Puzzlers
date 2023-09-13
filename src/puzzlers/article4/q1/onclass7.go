package main

import (
	"flag"
	"fmt"
)

func main() {
	// 变量声明
	// 方式1：显式声明
	// var name string
	// flag.StringVar(&name, "name", "everybody", "The greeting object")

	// flag.Parse()
	// fmt.Printf("Hello, %s!\n", name)

	// 方式2：隐式类型推断，可以在任意位置使用
	// var name = flag.String("name", "everybody", "The greeting object")

	// flag.Parse()
	// fmt.Printf("Hello, %v!\n", name)

	// 方式3：短类型声明，只可以在函数内部使用
	name := flag.String("name", "everybody", "The greeting object")

	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)

}
