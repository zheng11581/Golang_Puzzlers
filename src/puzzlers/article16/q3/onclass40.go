package main

import (
	"fmt"
	"sync/atomic"
	"time"
	// "sync/atomic"
	// "time"
)

func main() {
	// var count uint32
	// trigger := func(i uint32, fn func()) {
	// 	for {
	// 		if n := atomic.LoadUint32(&count); n == i {
	// 			fn()
	// 			atomic.AddUint32(&count, 1)
	// 			break
	// 		}
	// 		time.Sleep(time.Nanosecond)
	// 	}
	// }

	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	var num uint32 = 100
	for i := uint32(0); i < num; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}

	trigger(num, func() {})

}
