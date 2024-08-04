package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "123456"
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hash))

	errPass := bcrypt.CompareHashAndPassword(hash, []byte("123456"))

	if errPass != nil {
		fmt.Println("error pass")
	} else {
		fmt.Println("true pass")
	}
}
