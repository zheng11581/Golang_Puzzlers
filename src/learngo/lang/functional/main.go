package main

import "fmt"

func main() {
	//fmt.Println(add(1, 2))
	fmt.Println(adder()(1, 2))
	fmt.Println(adder2()(1, 2))
}

func add(a, b int) int {
	return a + b
}

func adder() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}

type myAdder func(int, int) int

func adder2() myAdder {
	return adder()
}
