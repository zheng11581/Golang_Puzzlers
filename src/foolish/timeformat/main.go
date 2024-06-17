package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	format := "02/Jan/2006:15:04:05 -0700"
	parse, err := time.Parse(format, "17/Jun/2024:06:50:26 +0000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parse)
}
