package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{})

	m1 := msgGen("producer1", done)
	for i := 0; i < 5; i++ {
		if msg, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println(msg)
		} else {
			fmt.Println("timeout")
		}
	}
	// 通知任务退出
	done <- struct{}{}
	<-done

	//shutdown channel
	//shutdown := make(chan os.Signal, 1)
	//select {}
	//signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	//<-shutdown

}

// 1. 生成器：服务、任务
func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleaning done")
				done <- struct{}{}
				return
			}
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

// 非阻塞等待
func nonBlockWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

// 超时等待
func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(2000 * time.Millisecond):
		return "", false
	}
}
