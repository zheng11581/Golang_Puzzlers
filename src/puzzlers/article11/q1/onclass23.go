package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// {
	// 	// 示例1：什么是单向通道？useless ?
	// 	uselessChan1 := make(chan<- int, 1) // 发送通道
	// 	uselessChan2 := make(<-chan int, 1) // 接收通道

	// 	// 打印的是代表两个通道指针的16进制
	// 	fmt.Printf("The useless channels: %v, %v\n", uselessChan1, uselessChan2)
	// }

	// {
	// 	// 示例2：单项通道作用
	// 	// 约束函数的定义
	// 	ch1 := make(chan int, 1)

	// 	go func() {
	// 		SendInt(ch1)
	// 	}()

	// 	for elem := range ch1 {
	// 		fmt.Printf("The value from channel: %d\n", elem)
	// 	}

	// }

	{
		// 示例2：单项通道作用
		// 约束方法的实现
		n := MyNotifier{}
		ch1 := make(chan int, 1)
		n.SendInt(ch1)
		fmt.Printf("The Send value is: %v\n", <-ch1)
		ch2 := n.GetIntChan(5)
		for elem := range ch2 {
			fmt.Printf("The Receive value is: %v\n", elem)
		}

	}
}

// func SendInt(ch1 chan<- int) {
// 	ch1 <- rand.Intn(1000)
// 	close(ch1) // 发送操作的地方结束一定要关闭通道
// }

type Notifier interface {
	SendInt(ch chan<- int)
	GetIntChan() <-chan int
}

type MyNotifier struct {
}

func (n MyNotifier) SendInt(ch1 chan<- int) {
	ch1 <- rand.Intn(1000)
	close(ch1) // 发送操作的地方结束一定要关闭通道
}

func (n MyNotifier) GetIntChan(length int) <-chan int {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // 发送操作的地方结束一定要关闭通道
	return ch
}
