package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "./files/abc.jfif"
	destFile := "./files/xyz.jfif"

	total, err := copyFile(srcFile, destFile)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Copy ok:", total)
	}
}

func copyFile(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}
