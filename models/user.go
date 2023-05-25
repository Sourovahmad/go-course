package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error){
	// 
	return nil, fmt.Errorf("error while %v")
}