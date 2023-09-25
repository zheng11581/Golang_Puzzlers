package main

import "fmt"

// 函数类型是一等数据类型
type Printer func(contents string) (n int, err error)

func printToStd(contents string) (n int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer

	// 可以作为值给变量赋值
	p = printToStd
	p("something")

	// 可以作为值进行类型转换、类型断言
	_, ok := interface{}(p).(Printer)
	fmt.Printf("The variable p's type is as same as Printer: %v\n", ok)
}
