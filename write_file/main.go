package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("I learn Golang!")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes\n", byteWritten)

	bs := []byte("Hello world")
	err = os.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
