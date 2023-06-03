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

func (us *UserService) Authenticate(email, password string) (*User, error) {

	requestEmail := strings.ToLower(email)

	loginuser := User{
		Email: requestEmail,
	}

	// get the user from databse.

	row := us.DB.QueryRow(`
		SELECT id,password FROM users WHERE email=$1
	`, requestEmail)

	err := row.Scan(&loginuser.Id, &loginuser.Password)

	if err != nil {
		return nil, fmt.Errorf("error on searching data using that email : %w", err)
	}

	// comparing the hash with the request hash

	compareError := bcrypt.CompareHashAndPassword([]byte(loginuser.Password), []byte(password))

	if compareError != nil {
		return nil, fmt.Errorf("the : %v  password is incorrect ", password)

	}

	return &loginuser, nil

}
