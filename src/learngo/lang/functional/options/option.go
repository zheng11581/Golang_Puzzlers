package main

import "fmt"

func main() {
	server := NewServer(WithPort(8090), WithTimeout(1000))
	fmt.Printf("Server Port: %d, Timeout: %d\n", server.Port, server.Timeout)
}

// Server 1. 定义 Server结构体
type Server struct {
	Port    int
	Timeout int
}

// Option 2. 定义 Option
type Option func(server *Server)

// WithTimeout 定义 Option函数
func WithTimeout(timeout int) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

// WithPort 定义 Option函数
func WithPort(port int) Option {
	return func(server *Server) {
		server.Port = port
	}
}

// NewServer Option模式来创建 Server实例
func NewServer(options ...Option) *Server {
	server := &Server{}
	for _, option := range options {
		option(server)
	}
	return server
}
