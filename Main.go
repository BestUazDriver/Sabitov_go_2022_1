package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello Bold!")
	file, err := os.OpenFile("hello.bin", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)
}
