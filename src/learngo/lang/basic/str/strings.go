package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "我爱北京天安门"

	// byte
	fmt.Println("Print as bytes")
	fmt.Println("字节数=", len(s))
	for i, b := range []byte(s) {
		fmt.Printf("(%d %X) ", i, b)
	}
	fmt.Println()

	// rune
	fmt.Println("Print as unicode")
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// char
	fmt.Println("Print as char with utf8.decode")
	bytes := []byte(s)
	fmt.Println("RuneCount=", utf8.RuneCount(bytes))
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	fmt.Println("Print as char simple")
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

	fmt.Println("strings 包")
	s = "我 我我 不知道"
	fmt.Printf("Fields=%v\n", strings.Fields(s))
	fmt.Printf("Spilt=%v\n", strings.Split(s, " "))
	fmt.Printf("Fields=%v\n", strings.Join(strings.Split(s, " "), ","))

	fmt.Printf("Contains=%v\n", strings.Contains(s, "我"))

}
