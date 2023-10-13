package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	num := uint32(18)
	fmt.Printf("The number: %d\n", num)
	delta := int32(-3)
	atomic.AddUint32(&num, uint32(delta))
	fmt.Printf("The number: %d\n", num)
	atomic.AddUint32(&num, uint32(-(-3)-1))
	fmt.Printf("The number: %d\n", num)

	fmt.Printf("The two's complement of %d: %b\n", delta, uint32(delta))
	fmt.Printf("The equivalent: %b\n", ^uint32(-(-3)-1)) // 与-3的补码相同。
}
