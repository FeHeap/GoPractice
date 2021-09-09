package main

import (
	"fmt"
	"os"
)

func main() {
	// fopen "r"
	fileName4 := "./test1/abc.txt"
	file2, err := os.Open(fileName4)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s is opened! %v\n", fileName4, file2)
	}

	var s []byte
	s = append(s, 'a')
	fmt.Println(s)
	_, err = file2.Write(s)
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	fileName5 := "./test1/abc.txt"
	file3, err := os.OpenFile(fileName5, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s is opened! %v\n", fileName5, file3)
	}

	file2.Close()
	file3.Close()
}