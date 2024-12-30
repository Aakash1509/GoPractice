package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	msg := "This is a demo file"

	file, err := os.Create("./myfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, msg)

	checkNilError(err)
	fmt.Println(length)

	defer file.Close()

	readFile("./myfile.txt")

}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)

	checkNilError(err)
	fmt.Println(string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
