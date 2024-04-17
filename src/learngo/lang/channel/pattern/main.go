package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	m1 := msgGen("producer1")
	m2 := msgGen("producer2")
	m := finIn(m1, m2)
	for {
		fmt.Println(<-m)
	}

}

// 1. 生成器：服务、任务
func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
		}
	}()
	return c
}

// 2. 同时等待多个服务、任务，不明确有多少个服务、任务
func finIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(ch chan string) {
			for {
				c <- <-ch
			}
		}(ch)
	}
	return c
}

// 2. 同时等待多个服务、任务 select，明确有多少个服务、任务
func finInBySelect(c1 chan string, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}

	}()
	return c
}
