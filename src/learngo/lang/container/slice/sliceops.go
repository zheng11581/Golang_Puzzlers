package main

import "fmt"

func updateSlice(s []int, i int, v int) {
	s[i] = v
}

func printSlice(s []int) {
	fmt.Printf("%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}
