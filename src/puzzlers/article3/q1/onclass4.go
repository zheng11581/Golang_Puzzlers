package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everybody", "The greeting object!")
}

func main() {
	flag.Parse()
	hello(name)
	// fmt.Printf("Hello, %s!\n", name)
}
