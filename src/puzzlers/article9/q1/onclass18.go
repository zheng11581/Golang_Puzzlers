package main

func main() {
	// 示例1
	// var badMap1 = map[[]int]int{} // 字典的键的类型不能是函数、字典、切片，否则编译错误
	// _ = badMap1

	// 示例2
	// var badMap2 = map[interface{}]int{
	// 	true:     1,
	// 	[]int{2}: 2,	// 如果键类型是 interface{}，实际的键值 也不能是函数、字典、切片，否则报 panic
	// 	"3":      3,
	// }
	// _ = badMap2

	// 示例3
	// var badMap3 map[[2][]string]int // 如果键类型是数组，数组的元素也不能是函数、字典、切片，编译错误
	// _ = badMap3

	// 示例4
	// type BadKey1 struct {
	// 	slice []string

	// }
	// var badMap4 map[BadKey1]int // 如果键类型是结构体，结构体的字段不能是函数、字典、切片，编译错误
	// _ = badMap4

	// 示例5
	// var badMap5 map[[2][3][4][]string]int // 不合法字段多深都会被揪出来
	// _ = badMap5

	// 示例6
	// type BadKey2Field1 struct {
	// 	slice []string
	// }
	// type BadKay2 struct {
	// 	field BadKey2Field1
	// }
	// var badMap6 map[BadKay2]int // 不合法字段多深都会被揪出来
	// _ = badMap6

}
