package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	// 定义多个通道
	intChannels := []chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	// 随机选取一个通道，并发送数据
	index := rand.Intn(3)
	intChannels[index] <- index

	// 哪个通道能接收到数据就执行哪个打印
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidateis selected")
	case <-intChannels[1]:
		fmt.Println("The secend candidate is selected")
	case <-intChannels[2]:
		fmt.Println("The third candidata is selected")
	default:
		fmt.Println("No candidate is selected")

	}

}

func example2() {
	intChan := make(chan int, 1)

	time.AfterFunc(time.Second, func() {
		close(intChan)
	})

	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The channel is closed")
			break
		}
		fmt.Println("The channel is not closed")
	}

}
