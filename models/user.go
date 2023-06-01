package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {

	requestEmail := strings.ToLower(email)

	hasedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("error hashing the password", password)
		return nil, err
	}

	paswordHash := string(hasedByte)

	user := User{
		Email:    requestEmail,
		Password: paswordHash,
	}

	row := us.DB.QueryRow(`
	INSERT INTO users (email, password)
	VALUES ($1,$2) RETURNING id`, requestEmail, paswordHash)

	scanError := row.Scan(&user.Id)

	if scanError != nil {
		fmt.Println("error scanning the Id", scanError)
		return nil, err
	}

	return &user, nil
}
