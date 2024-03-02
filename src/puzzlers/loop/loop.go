package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(convToBin(13)) // 1011 -> 1101
	printFile("abc.txt")
}

// 初始条件可以省略（类似 while）
func convToBin(n int) string {
	var result = ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 初始条件、递增条件可以省略（类似 while）
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("The file is not found")
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 初始条件、递增条件、结束条件都可以省略（死循环）
// func forever() {
// 	for {
// 		fmt.Println("abc")
// 	}
// }
