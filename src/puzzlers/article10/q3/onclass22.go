package main

import "fmt"

func main() {

	ch1 := make(chan int, 2)
	// 发送操作：生产者
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %d\n", i)
			ch1 <- i
		}
		fmt.Printf("Sender: close the channel %v\n", ch1)
		close(ch1)
	}()

	// 接收操作：消费者
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Printf("Receiver: channel closed")
			break
		}
		fmt.Printf("Receiver: receiving element %d\n", elem)
	}

}
