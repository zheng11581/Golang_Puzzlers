package main

import (
	"fmt"
	// "time"
)

func main() {
	num := 10
	sign := make(chan struct{}, num)
	for i := 0; i < 10; i++ {
		go func(i int) { // go函数执行会明显滞后于go语句
			fmt.Println(i)
			sign <- struct{}{}
		}(i)
	}

	// 方法1 sleep阻塞主goroutin，等待go函数执行完
	// time.Sleep(time.Millisecond * 500)

	// 方法2 channel阻塞主goroutin，等待go函数执行完
	for j := 0; j < 10; j++ {
		<-sign
	}

}
