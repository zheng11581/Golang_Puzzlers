package main

import (
	"fmt"
	"math"
)

func variableConvert() {
	var (
		a, b int = 3, 4
	)
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	variableConvert()
}
