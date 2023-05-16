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
	
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			roll INT,
			description TEXT
		);
	
	`)

	if err != nil {
		panic(err)
	}

	fmt.Println("table has been created")

}
