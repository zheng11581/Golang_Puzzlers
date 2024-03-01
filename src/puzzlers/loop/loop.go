package main

import "strconv"

func main() {
	println(convToBin(13)) // 1011 -> 1101
}

func convToBin(n int) string {
	var result = ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
