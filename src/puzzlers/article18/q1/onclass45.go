package main

import (
	"fmt"
	"os"
	"os/exec"
)

func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func main() {
	// 示例1
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("unexpacted error: %s\n", err)
		return
	}
	r.Close() // 人为制造 *os.PathError 类型错误
	_, err = w.Write([]byte("hi"))
	uError := underlyingError(err)
	fmt.Printf("underlying error: %s (type: %T)\n", uError, uError)
	r.Close()

	// 示例2
	paths := []string{
		os.Args[0], //
		"/it/must/not/exist",
		os.DevNull,
	}

	printError := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}
		err = underlyingError(err)
		switch err {
		case os.ErrClosed:
			fmt.Printf("error ")
		}
	}

}
