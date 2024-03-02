package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	if result, err := eval(3, 4, "+"); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println(result)
	}

	// 4.2 传递的函数参数也可以是匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 3, 5))

}

// 1. 函数签名：函数名+返回值
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported op: %s", op)

	}
}

// 2. 函数返回值可以是多个
// 3. 函数名可以有自己的名字（不建议）
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 4. 函数是一等公民：可以作为参数传递、返回值返回等
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 4.1 传递的函数参数可以有名字
func Pow(a, b int) int {
	result := math.Pow(float64(a), float64(b))
	return int(result)
}

// 函数的可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
