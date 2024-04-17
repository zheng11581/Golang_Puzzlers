package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	mu    sync.Mutex
}

func (a *atomicInt) increment() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go a.increment()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
