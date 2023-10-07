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
			fmt.Printf("error(closed)[%d]: %s\n", i, err)
		case os.ErrInvalid:
			fmt.Printf("error(invalid)[%d]: %s\n", i, err)
		case os.ErrPermission:
			fmt.Printf("error(permission)[%d]: %s\n", i, err)
		}
	}

	var f *os.File
	var index int
	{
		index = 0
		f, err = os.Open(paths[index])
		if err != nil {
			fmt.Printf("unexpected error: %s\n", err)
			return
		}
		f.Close() // 人为制造 os.ErrClosed 错误
		_, err = f.Read([]byte{})
		printError(index, err)
	}

	{
		index = 1
		// 人为制造 os.ErrInvalid 错误
		f, _ = os.Open(paths[index])
		_, err = f.Stat()
		printError(index, err)

	}

	{
		index = 2
		// 人为制造潜在错误为 os.ErrPermission 的错误。
		_, err = exec.LookPath(paths[index])
		printError(index, err)
	}

	if f != nil {
		f.Close()
	}
	fmt.Println()

}
