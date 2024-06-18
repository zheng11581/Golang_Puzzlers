package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	layout := "02/Jan/2006:15:04:05 -0700"

	parse, err := time.Parse(layout, "17/Jun/2024:06:50:26 +0000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parse string to time: %v\n", parse)

	format := time.Now().Format(layout)
	fmt.Printf("Format time to string: %v\n", format)
}
