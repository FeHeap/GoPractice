package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Mkdir
	fileName1 := "./test1"
	err := os.Mkdir(fileName1, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("Directory %s is created!\n", fileName1)
	}

	// os.MkdirALL
	fileName2 := "./test2/abc/xyz"
	err = os.MkdirAll(fileName2, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("Directory %s is created!\n", fileName2)
	}

	// os.Create
	fileName3 := "./test1/abc.txt"
	file1, err := os.Create(fileName3)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("File %s is created! %v\n", fileName3, file1)
	}
}