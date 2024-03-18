package main

import (
	"fmt"
	"time"
)

func main() {
	expireAt := time.Date(1988, 12, 21, 12, 30, 33, 0, time.Local).Unix()
	now := time.Now().Unix()
	fmt.Printf("1988-12-21 12:30:33: %d\n", expireAt)
	fmt.Printf("Now: %d\n", now)
	fmt.Println(expireAt < now)
}
