package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() { // go 函数执行会明显滞后于go 语句
			fmt.Println(i)
		}()
	}
}
