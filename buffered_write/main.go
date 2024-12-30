package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Opening the file for writing
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99} //While writing in a file , ASCII value will be written i.e abc

	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bytes written to buffer(not file): %d", bytesWritten)

	//checking available buffer
	bytesAvailable := bufferedWriter.Available()
	log.Printf("bytes available: %d", bytesAvailable)

	//Writing a string in buffer
	bytesWritten, err = bufferedWriter.WriteString("\nJust a random thing")

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// checking how much data is stored in buffer, just  waiting to be written to disk
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// The bytes have been written to buffer, not yet to file.
	// Writing from buffer to file.
	bufferedWriter.Flush()

}
