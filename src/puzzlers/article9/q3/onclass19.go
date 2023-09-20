package main

import "fmt"

func main() {
	// m := make(map[string]int)
	// fmt.Printf("m: %v (len: %d)\n", m, len(m))
	// m["1"] = 1
	// m["2"] = 2
	// m["3"] = 3
	// fmt.Printf("m: %v (len: %d)\n", m, len(m))

	var m map[string]int // 仅仅声明，字典、元素是“零值”，除了添加键 - 元素对，我们在一个值为nil的字典上做任何操作都不会引起错误
	fmt.Printf("m is nil: %v (len: %d)\n", m == nil, len(m))
	key := "two"
	elem, ok := m[key]
	fmt.Printf("The element pair with key %q in nil map: %d (%v)\n", key, elem, ok) // 对 nil 字典取值，elem 是 0，ok 是 false

	fmt.Printf("The length of nil map: %d\n", len(m))

	fmt.Printf("Delete the key-element pair by key %q\n", key) // 不会引发 panic
	delete(m, key)

	fmt.Printf("Add a key-element pair to nil map\n") // 会引发 panic
	m["two"] = 2
	fmt.Printf("The element with key (%s)：%d\n", key, m[key])
}
