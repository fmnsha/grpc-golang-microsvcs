package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/copier"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u *User) Test() {
	fmt.Println("test ")
}

func main() {

	u := &User{
		Name:  "feras",
		Age:   40,
		Email: "fmnsha@gmail.com",
	}
	c, err := DeepCopy(u)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

}

func DeepCopy[T any](source *T) (*T, error) {
	var target T

	if err := copier.Copy(&target, source); err != nil {
		return nil, err
	}

	return &target, nil
}
