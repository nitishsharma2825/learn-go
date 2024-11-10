package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s = s + sep + os.Args[i]
		sep = " "
	}
	s += "\n"
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	// strings.Join(os.Args[1:], " ")
	fmt.Printf(s)
}
