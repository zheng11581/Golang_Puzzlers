package main

import (
	"fmt"
	"puzzlers/article10/q0/lib"
)

func main() {
	ch1 := make(chan lib.Man, 3)
	m1 := lib.Man{
		Name:    "郑海成",
		Strong:  false,
		Rich:    false,
		Married: true,
	}
	ch1 <- m1
	m2 := lib.Man{
		Name:    "钢铁侠",
		Strong:  true,
		Rich:    true,
		Married: true,
	}
	ch1 <- m2
	m3 := lib.Man{
		Name:   "蝙蝠侠",
		Strong: false,
		Rich:   true,
	}
	ch1 <- m3

	nm1 := <-ch1
	fmt.Printf("Hero name: %q, Power: %v\n", nm1.Name, nm1.Strong)
}
