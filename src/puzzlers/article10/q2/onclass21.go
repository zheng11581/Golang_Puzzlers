package main

func main() {
	// ch1 := make(chan int, 1)
	// go func() {
	// 	ch1 <- 1
	// 	ch1 <- 2
	// }()
	// go func() {
	// 	fmt.Printf("ch1: %d\n", <-ch1)
	// 	fmt.Printf("ch1: %d\n", <-ch1)
	// }()
	// time.Sleep(time.Second * 10)

	// 示例1
	ch1 := make(chan int, 1)
	ch1 <- 1
	// ch1 <- 2 // 通道已满，这里会被阻塞

	// 示例2
	ch2 := make(chan int, 1)
	// elem, ok := <-ch2 //通道已空，这里会被阻塞
	// _, _ = elem, ok
	ch2 <- 1

	// 示例3
	var ch3 chan int
	// ch3 <- 1 // 通道的值为 nil，发送操作会被阻塞
	_ = <-ch3 // 通道的值为 nil，接收操作会被阻塞
}
