package main

import "fmt"

func main() {
	// 示例1
	// value3 := [...]int{0, 1, 2, 3, 4, 5, 6}
	// switch value3[0] {
	// case 0, 1, 2: // 编译无法通过，case 子表达式:2不能重复
	// 	fmt.Println("0 or 1 or 2")
	// case 2, 3, 4: // 编译无法通过，case 子表达式:2 4不能重复
	// 	fmt.Println("2 or 3 or 4")
	// case 4, 5, 6: // 编译无法通过，case 子表达式:4不能重复
	// 	fmt.Println("4 or 5 or 6")
	// }

	// 示例2
	value4 := [...]int{0, 1, 2, 3, 4, 5, 6}
	switch value4[0] {
	case value4[0], value4[1], value4[2]: // 可以绕过编译检查
		fmt.Println("0 or 1 or 2")
	case value4[2], value4[3], value4[4]:
		fmt.Println("2 or 3 or 4")
	case value4[4], value4[5], value4[6]:
		fmt.Println("4 or 5 or 6")
	}
}
