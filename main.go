package main

import (
	"fmt"
	"os"

	qrCode "github.com/skip2/go-qrcode"
)

const (
	MAX_SIZE          = 1024
	QR_CODED_PNG_SIZE = 256
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

func createQrCode(data string, destination string) {
	err := qrCode.WriteFile(data, qrCode.Medium, QR_CODED_PNG_SIZE, destination)
	checkError(err)
}

func main() {
	data := readFile("input.txt")
	createQrCode(data, "output.png")
	fmt.Println("QR Code created successfully!")
}
