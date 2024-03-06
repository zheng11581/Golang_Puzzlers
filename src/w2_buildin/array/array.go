package main

import "fmt"

func main() {
	// 1. 数组的定义
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 3, 5, 7, 9}
	var arr4 [2][2]int
	fmt.Println(arr1, arr2, arr3, arr4)

	// 2. 数组的遍历
	// 2.1 第一种
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 2.2 第二种：可以同时遍历index和value
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("printArrayValue(arr1)")
	printArrayValue(arr1)
	fmt.Println("printArrayValue(arr3)")
	printArrayValue(arr3)
	fmt.Println(arr1, arr3)

	fmt.Println("printArrayRef(arr1)")
	printArrayRef(&arr1)
	fmt.Println("printArrayRef(arr3)")
	printArrayRef(&arr3)
	fmt.Println(arr1, arr3)

}

// 3. 数组是值类型
// 不会改变函数外面的数组
func printArrayValue(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayRef(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
