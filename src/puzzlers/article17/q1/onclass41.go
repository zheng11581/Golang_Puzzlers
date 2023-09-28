package main

import "fmt"

func main() {
	// 示例1
	number1 := []int{1, 2, 3, 4, 5, 6}
	for i := range number1 {
		if i == 3 {
			number1[i] |= i
		}
	}
	fmt.Println(number1)
	fmt.Println()

	// 示例2
	number2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(number1) - 1
	for i, e := range number2 {
		if maxIndex2 == i {
			number2[0] += e
		} else {
			number2[i+1] += e
		}
	}
	fmt.Println(number2)
	fmt.Println()

	// 示例3
	number3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(number1) - 1
	for i, e := range number3 {
		if maxIndex3 == i {
			number3[0] += e
		} else {
			number3[i+1] += e
		}
	}
	fmt.Println(number3)
	fmt.Println()
}
