package main

import (
	"flag"
	"os"
	fg "puzzlers/article3/q5/flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fg.Hello(os.Stdout, name)
}
