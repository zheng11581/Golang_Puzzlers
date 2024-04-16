package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	var c = make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func worker(id int, c chan int) {
	for v := range c {
		// 消费者速度慢
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d received %d\n", id, v)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	// producer
	var c1, c2 = generator(), generator()
	// consumer
	var w = createWorker(0)
	var values []int

	// 10s之后向chan中发送一个，总的执行时间
	tm := time.After(time.Second * 10)
	// 每隔1s向chan发送一个
	tk := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}

		// default case 非阻塞
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout!")
		case <-tk:
			fmt.Println("queue len = ", len(values))
		case <-tm:
			fmt.Println("bye eye!")
			return
		}

	}
}
