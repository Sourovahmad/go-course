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

	// _, err = db.Exec(`
	// 	INSERT INTO users (name,email) VALUES ('sourov','sourfafssdffsdfsaov2@gmail.com');
	// `)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("data inserted into db")

	//  query in the database

	// id := 4

	// row := db.QueryRow(`

	// SELECT name, email
	// FROM users

	// WHERE id=$1;`, id)

	// var name, email string

	// err = row.Scan(&name, &email)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("the name and email is", name, email)

	// inserting fake data into database

	UserId := 2
	for i := 1; i <= 5; i++ {
		roll := i
		description := fmt.Sprintf("fake description #%d", i)

		_, err := db.Exec(`
			INSERT INTO students(user_id,roll,description) VALUES ($1,$2,$3);
		`, UserId, roll, description)

		if err != nil {
			panic(err)
		}

	}

	// --------------- query on multiple items --------------------

	type Student struct {
		Id          int
		UserId      int
		Roll        int
		Description string
	}
	var students []Student

	rows, rowerr := db.Query(`
		SELECT id,roll,description FROM students WHERE user_id=$1;
	`, UserId)

	if rowerr != nil {
		panic(rowerr)
	}

	defer rows.Close()

	// 	query the data and insert into student struct

	for rows.Next() {
		var student Student

		student.UserId = UserId
		err := rows.Scan(&student.Id, &student.Roll, &student.Description)

		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}

	fmt.Println(students)

}
