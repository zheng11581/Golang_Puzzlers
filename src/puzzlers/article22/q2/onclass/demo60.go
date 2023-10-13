package main

import (
	"fmt"
	"sync"
	"time"
)

// counter 代表计数器
type counter struct {
	count uint
	mu    sync.RWMutex
}

// number 返回计数器的计数
func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

// add 计数器加1，返回计数器计数
func (c *counter) add(increase uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count += increase
	return c.count
}

func main() {
	// c := counter{}
	// count(&c)
	redundantUnlock()
}

func count(c *counter) {
	sign := make(chan struct{}, 3)

	// 增加计数器计数
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()

	// 显示计数器计数
	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 200)
			fmt.Printf("The number in counter: %d [%d-%d]\n", c.number(), 1, j)
		}
	}()

	// 显示计数器计数
	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for k := 1; k <= 20; k++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Printf("The number in counter: %d [%d-%d]\n", c.number(), 2, k)
		}
	}()
	<-sign
	<-sign
	<-sign
}

func redundantUnlock() {
	var rwMu sync.RWMutex

	// 示例1。
	// rwMu.Unlock() // 这里会引发panic。

	// 示例2
	// rwMu.RUnlock() // 这里会引发panic。

	// 示例3
	rwMu.RLock()
	// rwMu.Unlock() // 这里会引发panic。
	rwMu.RUnlock()

	// 示例4
	rwMu.Lock()
	// rwMu.RUnlock() // 这里会引发panic。
	rwMu.Unlock()

}
