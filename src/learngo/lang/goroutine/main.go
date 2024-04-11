package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		//go func(i int) {
		//	fmt.Printf("Hello from %d\n", i)
		//}(i)

		go func(i int) {
			sum := 0
			for v := 0; v <= i; v++ {
				sum += v
			}
			fmt.Printf("0 + 1 + ... + %d = %d\n", i, sum)
		}(i)
	}
	time.Sleep(1 * time.Millisecond)

}

// 1. 非抢占式
// 2. 让出时机
// 3. 多个goroutine可能绑定到同一个线程
// 4. 调用函数加一个go关键字就可以
