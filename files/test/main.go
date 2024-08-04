package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	path := "client_flies"

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	filePath := filepath.Join(path, "myfile.txt")

	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(file)

}
