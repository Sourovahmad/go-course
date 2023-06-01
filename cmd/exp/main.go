package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgressConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
}

func (cfg PostgressConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SslMode)
}

func main() {

	config := PostgressConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "postgressDb",
		SslMode:  "disable",
	}

	db, err := sql.Open("pgx", config.String())
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to db")

	//  create tables

	_, err = db.Exec(`
	
			CREATE TABLE IF NOT EXISTS users(
				id SERIAL PRIMARY KEY,
				email TEXT UNIQUE NOT NULL,
				password TEXT NOT NULL
			);

	
	`)

	if err != nil {
		panic(err)
	}

	fmt.Println("table has been created")

	user := struct {
		Id       int
		Email    string
		Password string
	}{
		Email:    "sourov_tesing@gmail.com",
		Password: "4521254fasdfasdf4ff5as4df54asf24as",
	}

	row := db.QueryRow(`
	INSERT INTO users (email, password)
	VAlUES ($1,$2) RETURNING id`, user.Email, user.Password)

	scanError := row.Scan(&user.Id)

	if scanError != nil {
		fmt.Println("error while scanning the id", scanError)
	}

	fmt.Println("user has been created", user)

}
