package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("invalid command: %v\n", os.Args[1])
	}
}

func hash(password string) {
	hasedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("error hashing the password", password)
		return
	}


	
	fmt.Println(string(hasedByte))
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		fmt.Printf("the : %v  password is incorrect ", password)
		return
	}

	fmt.Println("password matched")
}
