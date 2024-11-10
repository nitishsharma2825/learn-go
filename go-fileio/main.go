package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	reader := io.Reader(os.Stdin)
	writer := io.Writer(os.Stdout)
	buffer := make([]byte, 1)
	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			break
		}
		bytesWritten, err := writer.Write(buffer)
		if err != nil {
			break
		}
		fmt.Printf("%v:%v", bytesRead, bytesWritten)
	}
}
