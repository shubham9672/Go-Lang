package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func fileHandling() {
	//Reading file using os.ReadFile
	data, err := os.ReadFile("file.txt")
	checkError(err)
	fmt.Println("--------File.txt Data-----")
	fmt.Println(string(data))

	//Reading file  using byte
	file, err := os.Open("file.txt")
	checkError(err)
	// reading first 6 bytes in file.txt
	//output [70 111 114 32 71 111]
	//For Go
	fileB1 := make([]byte, 6)
	fileRead, err := file.Read(fileB1)
	checkError(err)
	fmt.Println(fileB1)
	fmt.Println(string(fileB1[:fileRead]))
	// using seek
	fSeek, err := file.Seek(4, 0)
	checkError(err)
	fmt.Println(fSeek)
	// reading with buffer
	bufferRead := bufio.NewReader(file)
	brPeek, err := bufferRead.Peek(12)
	checkError(err)
	fmt.Println(string(brPeek))

	/* ------ Writing data into file-----*/
	file1, err := os.Create("fileWrite.txt")
	checkError(err)
	fmt.Println("Enter Text To write into file:")
	// input: why learn go lang?
	inputReader := bufio.NewReader(os.Stdin)
	text, _ := inputReader.ReadString('\n')
	file1.WriteString(text)
	// ouput: written  into fileWrite.txt
	fmt.Print("Written into ", file1.Name())
	//Writing file using bytes read from input.txt using fileb1
	file2, err := os.Create("fileByteWrite.txt")
	checkError(err)
	file2.Write(fileB1)
	fmt.Println("Written into ", file2.Name(), " Using byte array read from file.txt")
	// writing string using byte array using os
	byteText := []byte("For Go Lang File Handling")
	err1 := os.WriteFile("file.txt", byteText, 0644)
	checkError(err1)
	file.Close()
	file1.Close()
	file2.Close()
	/*---------Output-------*/
	/*
		--------File.txt Data-----
		[70 111 114 32 71 111]
		For Go
		4
		Go Lang File
		Enter Text To write into file:
		Why learn go lang?
		Written into fileWrite.txt
		Written into fileByteWrite.txt Using byte array read from file.txt
	*/
}
