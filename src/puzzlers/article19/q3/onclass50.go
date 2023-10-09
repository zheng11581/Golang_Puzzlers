package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter funtion main.")

	defer func() {
		fmt.Println("Enter funtion defer.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exist funtion defer.")
	}()

	// recover() 错误使用方法，此时没有 panic
	fmt.Printf("no panic: %s\n", recover())

	// 用 panic() 主动引发panic
	panic(errors.New("something wrong"))

	// recover() 错误使用方法，没有机会执行
	fmt.Printf("panic: %s", recover())

	fmt.Println("Exist funtion main.")
}
