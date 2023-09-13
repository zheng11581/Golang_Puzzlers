package main

import (
	"flag"
	lib "puzzlers/article3/q2/onclasslib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everybody", "The greeting object")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}
