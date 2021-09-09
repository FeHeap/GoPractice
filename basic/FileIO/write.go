package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./files/abc.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("Error: open file!", err.Error())
	} else {
		n, err := file.Write([]byte("abcde12345"))
		if err != nil {
			fmt.Println("Error: Write file!", err.Error())
		} else {
			fmt.Println("OK:", n)
		}

		n, err = file.WriteString("中國字")
		if err != nil {
			fmt.Println("Error: Write file!", err.Error())
		} else {
			fmt.Println("OK:", n)
		}
	}
}