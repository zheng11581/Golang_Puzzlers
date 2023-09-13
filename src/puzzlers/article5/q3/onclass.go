package main

import (
	"fmt"
	// main 包作用域[1]
	// . "os"
)

// main 包作用域，不可以和上面位置[1]重复
// const O_RDONLY = 1

func main() {
	// main 函数作用域，可以和上面位置[1]重复
	const O_RDONLY = 1
	fmt.Printf("The const is: %d\n", O_RDONLY)
}
