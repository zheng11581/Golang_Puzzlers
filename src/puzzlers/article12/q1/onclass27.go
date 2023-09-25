package main

import (
	"errors"
	"fmt"
)

type operate func(x int, y int) int

func main() {
	// 方案1
	x, y := 10, 20
	op := func(x int, y int) (result int) {
		return x + y
	}
	result, err := calculator(x, y, op)
	fmt.Printf("The result is %d (error: %v)\n", result, err)
	result, err = calculator(x, y, nil)
	fmt.Printf("The result is %d (error: %v)\n", result, err)

	// 方案2
	x, y = 128, 256
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result is %d (error: %v)\n", result, err)

}

// 方案1：把函数作为参数传入
func calculator(x, y int, op operate) (n int, err error) {
	if op == nil {
		return 0, errors.New("invalide operation")
	}
	result := op(x, y)
	return result, nil
}

// 方案2：把函数作为结果返回
type calculateFunc func(x, y int) (result int, err error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (result int, err error) {
		if op == nil {
			return 0, errors.New("invalide operation")
		}
		return op(x, y), nil
	}
}
