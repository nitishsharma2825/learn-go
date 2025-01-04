package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, _ := rpc.Dial("tcp", "localhost:1234")
	args := Args{A: 7, B: 11}
	var reply int

	err := client.Call("Calculator.Add", &args, &reply)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", reply)
	}
}
