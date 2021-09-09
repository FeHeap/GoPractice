package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./files/blockchain.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: open file!", err.Error())
	} else {
		bs := make([]byte, 1024*8, 1024*8)
		n := -1
		for {
			n, err = file.Read(bs)
			if n == 0 || err == io.EOF {
				fmt.Println("Read over!")
				break
			}
			fmt.Println(string(bs[:n]))
		}
	}
	file.Close()
}
