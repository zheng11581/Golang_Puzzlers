package main

import (
	"flag"
	// inlib "puzzlers/article3/q4/onclasslib/internal"
	lib "puzzlers/article3/q4/onclasslib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everybody", "The greeting object")
}

func main() {
	flag.Parse()
	// 无法引用
	// inlib.Hello(name)
	// 可以引用
	lib.Hello(name)
}
