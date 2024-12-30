package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("test.txt") //default permissions will be : -rw-rw-r--
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Created file test.txt")
	}

	//err = os.Truncate("test.txt", 0) //Only 20 bytes will be there rest will be deleted , 0 means completely empty the file
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = newFile.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//file.Close()

	//File information
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("test.txt")
	p := fmt.Println
	p("File Name: ", fileInfo.Name())
	p("File Size: ", fileInfo.Size())
	p("Permissions: ", fileInfo.Mode())
	p("Last Modified: ", fileInfo.ModTime())
	p("IsDirectory: ", fileInfo.IsDir())

	fileInfo, err = os.Stat("a.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
	}

	//oldPath := "test.txt"
	//newPath := "test1.txt"
	//err = os.Rename(oldPath, newPath)
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = os.Remove("test1.txt")
	if err != nil {
		log.Fatal(err)
	}

}
