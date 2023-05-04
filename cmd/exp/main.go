package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateOrg()
	fmt.Println(err)
}

func ConnectDb() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := ConnectDb()
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	return nil
}

func CreateOrg() error {
	orgError := CreateUser()
	if orgError != nil {
		return fmt.Errorf("org create error: %w", orgError)
	}
	return nil
}
