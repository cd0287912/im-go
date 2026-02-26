package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := Server{
		IP:   ip,
		Port: port,
	}
	return &server
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("已连接")
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("listener is err", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err", err)
		}
		go this.Handler(conn)
	}
}
