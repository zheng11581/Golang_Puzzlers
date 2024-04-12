package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, worker worker) {
	for v := range worker.in {
		fmt.Printf("Worker %d received %c\n", id, v)
		worker.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		make(chan int),
		func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	// wait for all of them is done
	//for _, worker := range workers {
	//	<-worker.done
	//}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	// wait for all of them is done
	//for _, worker := range workers {
	//	<-worker.done
	//}
	wg.Wait()

	// 通过 waitGroup 去等待执行完
	// 通过 done chan 去通知执行完了，不用 sleep了
	// time.Sleep(1 * time.Millisecond)

}

func main() {
	chanDemo()
}
