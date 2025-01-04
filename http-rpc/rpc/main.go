package main

import (
	"net"
	"net/rpc"
)

type Args struct{ A, B int }

type Calculator int

func (c *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, _ := net.Listen("tcp", ":1234")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
