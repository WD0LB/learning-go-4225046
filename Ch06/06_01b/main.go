package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")
	fileName := "./myFile.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	checkError(err)
	length, err := io.WriteString(file, "Let's write in file")
	fmt.Println("Wrote a file with number of characters:", length)
	readFile(fileName)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}