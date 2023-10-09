package main

import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main() {
	flag.Parse()
	greeting, err := hello(name)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println(greeting, introduce())
}

func introduce() (introduce string) {
	introduce = fmt.Sprintf("Welcome to Golang column.")
	return
}

func hello(name string) (greeting string, err error) {
	if name == "" {
		err = errors.New("empty name")
		return
	}
	greeting = fmt.Sprintf("Hello, %s!", name)
	return
}
