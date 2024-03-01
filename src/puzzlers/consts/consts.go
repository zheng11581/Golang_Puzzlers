package main

import "fmt"

func main() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		pb
	)
	fmt.Println(b, kb, mb, gb, gb, pb)
}
