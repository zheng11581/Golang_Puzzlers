package main

import "fmt"

func main() {
	// 这个链表与该defer语句所属的函数是对应的，并且，它是先进后出（FILO）的，相当于一个栈。
	defer fmt.Println("first defer.")
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Printf("defer in for [%d].\n", i)
		}(i)
	}
	defer fmt.Println("last defer.")
}
