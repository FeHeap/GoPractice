package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	testReader()
	fmt.Println("-----")

	testWriter()
	fmt.Println("-----")

	reader := bufio.NewReader(strings.NewReader("abcdefg 1000phone blockchain ready go"),)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "q!" {
			break
		}
	}
}

func testReader() {
	fileName := "./files/blockchain.txt"
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	fmt.Printf("%T\n", reader)
	for {
		s, err := reader.ReadString('\n')
		fmt.Print(s)
		if err == io.EOF {
			fmt.Println("\nRead over")
			break
		}
	}
	file.Close()
}

func testWriter() {
	filename1 := "./files/xyz.jfif"
	file1, _ := os.Open(filename1)
	reader := bufio.NewReader(file1)
	filename2 := "./files/pqr.jfif"
	file2, _ := os.OpenFile(filename2, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	writer := bufio.NewWriter(file2)
	for {
		bs, err := reader.ReadBytes(' ')
		writer.Write(bs)
		writer.Flush()
		if err == io.EOF {
			fmt.Println("Read over")
			break
		}
	}
	file1.Close()
	file2.Close()
}