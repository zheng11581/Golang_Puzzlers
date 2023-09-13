package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	// 这里对 err 进行了重生明
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("%d byte(s) were been writen.", n)
}
