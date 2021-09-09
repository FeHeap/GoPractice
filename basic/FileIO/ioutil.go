package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName1 := "./files/blockchain.txt"
	fileName2 := "./files/xyz.txt"
	dirName := "./"

	// ReadFile
	data, err := ioutil.ReadFile(fileName1)
	if err != nil {
		fmt.Println("Error: open file!", err.Error())
	} else {
		fmt.Println(string(data))
	}

	// WriteFile
	s1 := "Hello, Go ~~"
	err = ioutil.WriteFile(fileName2, []byte(s1), 0777)
	if err != nil {
		fmt.Println("Error: Write file!", err.Error())
	} else {
		fmt.Println("OK: write file!")
	}

	// Copy
	err = ioutil.WriteFile(fileName2, data, os.ModePerm)
	if err != nil {
		fmt.Println("Error: Copy file!", err.Error())
	} else {
		fmt.Println("OK: copy file!")
	}

	// Traverse directory
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error: Traverse directory!", err.Error())
	} else {
		for i, v := range fileInfos {
			fmt.Println(i, v.Name(), v.IsDir(), v.Size(), v.ModTime())
		}
	}

	// Create directory
	filename, err := ioutil.TempDir("./", "temp")
	if err != nil {
		fmt.Println("Error: create directory!", err.Error())
	} else {
		fmt.Println(filename)
	}

	// Create file
	file, err := ioutil.TempFile("./", "temp")
	if err != nil {
		fmt.Println("Error: create file!", err.Error())
	} else {
		file.WriteString("Content:" + file.Name())
	}
	file.Close()
}
