package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// mailbox 代表信箱
	// 0 信箱无信件，1 信箱有信件
	var mailbox int

	// mailLock 代表信箱的锁
	var lock sync.RWMutex

	// sendCond 代表发信的条件变量
	sendCond := sync.NewCond(&lock)

	// recvCond 代表收信的条件变量
	recvCond := sync.NewCond(lock.RLocker())

	// 收信、发信信号
	sign := make(chan struct{}, 3)
	max := 5

	// 开始发信
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox > 0 {
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}()

	// 开始收信
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has received.", j)
			lock.RUnlock()
			sendCond.Signal()

		}
	}()

	<-sign
	<-sign

}
