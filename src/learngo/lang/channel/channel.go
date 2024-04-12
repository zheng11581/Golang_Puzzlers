package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// 5. 用 range接收 channel，可以自动判断 channel结束
	for v := range c {
		fmt.Printf("Worker %d received %d\n", id, v)
	}
}

// 2. channel 是一等公民，可以作为参数和返回值传递
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	// 3.channel 可以声明成单向的，只能接受操作或者只能发送操作
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		// 1. channel 发送操作和接受操作应该在两个goroutine中成对使用，否则可能 deadlock
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(1 * time.Millisecond)
}

func bufferedChan() {
	// 4. buffer channel
	c := make(chan int, 2)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(1 * time.Millisecond)
}

func chanClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	// 6. 发送方 close，接收方处理（1）range（2）v, ok := <- c
	close(c)
	time.Sleep(1 * time.Millisecond)
}

func main() {
	fmt.Println("The first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	//bufferedChan()
	fmt.Println("Close channel")
	//chanClose()
}
