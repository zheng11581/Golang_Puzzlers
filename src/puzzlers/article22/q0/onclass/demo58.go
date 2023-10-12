package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
)

// protecting 用于指示是否使用互斥锁来保护数据写入。
// 若值等于0则表示不使用，若值大于0则表示使用。
// 改变该变量的值，然后多运行几次程序，并观察程序打印的内容。
var protecting uint = 1

func main() {
	// buffer 代表缓冲区。
	var buffer bytes.Buffer

	// max1 代表启用的goroutine的数量。
	// max2 代表每个goroutine需要写入的数据块的数量。
	// max3 代表每个数据块中需要有多少个重复的数字。
	const (
		max1 = 5
		max2 = 10
		max3 = 10
	)

	// mu 代表以下流程要使用的互斥锁。
	var mu sync.Mutex
	// sign 代表信号的通道。
	sign := make(chan struct{}, max1)

	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j < max2; j++ {
				// 准备数据。
				header := fmt.Sprintf("\n[id: %d, iteration: %d]", id, j)

				// 加锁
				data := fmt.Sprintf(" %d", id*j)
				if protecting > 0 {
					mu.Lock()
				}

				// 开始写入数据
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]", err, id)
				}
				for k := 1; k < max3; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]", err, id)
					}
				}

				// 解锁
				if protecting > 0 {
					mu.Unlock()
				}
			}

		}(i, &buffer)
	}

	for i := 1; i < max1; i++ {
		<-sign
	}

	data, err := io.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents: %s\n", data)
}
