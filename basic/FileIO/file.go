package main

import (
	"fmt"
	"os"
)

func main() {
	// Absolute
	path := "E:/Git/PythonPractice/WordCloud/MonikaOnly.jpg"
	printMessage(path)

	// Relative
	path = "../PerlPractice/basic/ch5_IO/tac.pl"
	printMessage(path)
}

func printMessage(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("Data type: %T\n", fileInfo)
		fmt.Println("File name:", fileInfo.Name())
		fmt.Println("If it's directory:", fileInfo.IsDir())
		fmt.Println("Size:", fileInfo.Size())
		fmt.Println("Authority:", fileInfo.Mode())
		fmt.Println("Last updated:", fileInfo.ModTime())
		fmt.Println()
	}
}