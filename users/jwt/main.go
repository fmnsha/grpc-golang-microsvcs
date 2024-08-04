package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserId string
	Role   string
}

func main() {

	key := []byte("mysecret") /* Load key from somewhere, for example a file */
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{UserId: "1", Role: "1"})
	s, err := t.SignedString(key)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	var claims UserClaims
	token, err := jwt.ParseWithClaims(s, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecret"), nil
	})

	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*UserClaims); ok {
		fmt.Println(claims.UserId, claims.Role, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}

}
