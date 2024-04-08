package main

import (
	"fmt"
	"os"
)

const (
	MAX_SIZE = 1024
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) string {
	file, err := os.Open(path)
	checkError(err)
	defer file.Close()

	data := make([]byte, MAX_SIZE)
	count, err := file.Read(data)
	checkError(err)

	return string(data[:count])
}

func main() {
	data := readFile("input.txt")
	fmt.Println(data)
}
