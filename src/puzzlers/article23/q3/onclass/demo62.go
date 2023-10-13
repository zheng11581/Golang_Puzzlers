package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// mailbox 代表信箱
	// 0 信箱无信件，1 信箱有信件
	var mailbox uint

	// mailLock 代表信箱的锁
	var lock sync.RWMutex

	// sendCond 代表发信的条件变量
	sendCond := sync.NewCond(&lock)

	// recvCond 代表收信的条件变量
	recvCond := sync.NewCond(lock.RLocker())

	send := func(id, index int) {
		lock.Lock()
		for mailbox > 0 {
			sendCond.Wait()
		}
		log.Printf("sender [%d-%d]: the mailbox is empty.", id, index)
		mailbox = 1
		log.Printf("sender [%d-%d]: the letter is sent.", id, index)
		lock.Unlock()
		recvCond.Broadcast()

	}

	recv := func(id, index int) {
		lock.RLock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		log.Printf("receiver [%d-%d]: the mailbox is full.", id, index)
		mailbox = 0
		log.Printf("receiver [%d-%d]: the letter is received.", id, index)
		lock.RUnlock()
		sendCond.Signal()
	}

	// 收信、发信信号
	sign := make(chan struct{}, 3)
	max := 6

	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(id, i)
		}
	}(0, max)

	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, j)
		}
	}(1, max/2)

	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= max; k++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, k)
		}
	}(2, max/2)

	<-sign
	<-sign
	<-sign

}
