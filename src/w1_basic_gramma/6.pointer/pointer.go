package main

import "fmt"

func main() {
	// 指针不能运算
	var a int = 2
	var b int = 20

	// *p 指向 a
	var p *int = &a
	*p = 3
	fmt.Printf("a 被修改为3: %d\n", a)

	// *p 指向 b
	p = &b
	*p = 23 // b 被修改为23
	fmt.Printf("a 值不变: %d, b 被修改为23: %d\n", a, b)

	swap(&a, &b)
	fmt.Printf("a: %d, b: %d\n", a, b)

}

func swap(a, b *int) {
	*b, *a = *a, *b
}

//  1. go 里面只有值传递：所传递的是该参数的副本，是复制了一份的，本质上不能认为是一个东西，指向的不是一个内存地址
// 非引用类型（值类型）：int，string，float，bool，数组和struct，传递的是副本的值
// 引用类型：指针，slice，map，channel，接口，函数
