package main

import (
	"fmt"
)

func main() {
	// 1. slice 可以从 array得来 [:]
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])

	// 2. slice 本身是没有数据的，是对array的引用（view），是引用类型（参数传递类似指针）
	fmt.Println("Before update: arr=", arr)
	updateSlice(arr[:], 0, 100)
	fmt.Println("After updated: arr=", arr)

	// 3. Reslice
	s1 := arr[2:6]
	fmt.Println("s1=", s1)
	s1 = s1[3:5]
	fmt.Println("s1=", s1)
	s1 = arr[:]
	fmt.Println("s1=", s1)

	// 4. Extends slice
	arr[0] = 0
	s1 = arr[2:6]
	fmt.Printf("s1=%v, lens(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := s1[3:5]
	fmt.Printf("s2=%v, lens(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	fmt.Println("Appending Slice")
	// 向slice添加元素时候如果超越cap，系统会分配更大的底层数组
	s3 := append(s2, 10)
	// s4、s5 no longer view arr.
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s5=", s5)
	fmt.Println("s2=", s2)
	fmt.Println("arr=", arr)

	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
		printSlice(s)
	}

	s11 := []int{0, 1, 2, 3, 4}
	printSlice(s11)

	s12 := make([]int, 10)
	printSlice(s12)

	s13 := make([]int, 10, 32)
	printSlice(s13)

	fmt.Println("Coping slice")
	copy(s12, s11)
	printSlice(s12)

	fmt.Println("Deleting from slice")
	s12 = append(s12[:4], s12[5:]...)
	printSlice(s12)

	fmt.Println("Poping from slice")
	head := s12[0]
	s12 = s12[1:]
	fmt.Println("Head of s12: ", head)
	printSlice(s12)

	fmt.Println("Droping from slice")
	tail := s12[len(s12)-1]
	s12 = s12[:len(s12)-1]
	fmt.Println("Tail of s12: ", tail)
	printSlice(s12)

	// s1 := make([]int, 5)
	// s2 := make([]int, 0)
	// fmt.Println(reflect.TypeOf(s1))
	// fmt.Println(reflect.TypeOf(s2))

}
