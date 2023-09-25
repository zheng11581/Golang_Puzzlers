package main

import (
	"fmt"
)

var channels = []chan int{
	nil,
	make(chan int, 1),
	nil,
}

var numbers = []int{1, 2, 3}

func main() {

	select {
	case getChan(0) <- getNumber(0): // nil <- 0，不满足选择条件
		fmt.Println("The first candidate case is selected!")
	case getChan(1) <- getNumber(1): // 非缓冲通道只有在收发双方都就绪的情况下才能传递元素值，否则就阻塞。
		fmt.Println("The secend candidate case is selected!")
	case getChan(2) <- getNumber(2): // nil <- 0，不满足选择条件
		fmt.Println("The third candidate case is selected!")
	default:
		fmt.Println("No candidate case is selected!")
	}

	for i := 0; i < 3; i++ {
		select {
		case channels[1] <- 1: // channel[1] 已经在上一个select 中满了，再发送阻塞，不满足选择条件
			fmt.Println("The candidate case is selected!")
		default:
			fmt.Println("No candidate case is selected!")
		}
	}

}

func getChan(n int) chan int {
	fmt.Printf("channels[%d]\n", n)
	return channels[n]
}

func getNumber(n int) int {
	fmt.Printf("numbers[%d]\n", n)
	return numbers[n]
}
