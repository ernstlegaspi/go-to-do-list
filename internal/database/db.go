package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBType struct {
	DB *sql.DB
}

func ConnectDB() (*DBType, error) {
	conn := "user=postgres dbname=postgres password=todolist sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error connecting to db...")
		return nil, err
	}

	return &DBType{
		DB: db,
	}, nil
}

func (d *DBType) CreateTables() error {
	if err := d.createUserTable(); err != nil {
		fmt.Println(err)
		fmt.Println("Can not create create list table.")
		return err
	}

	if err := d.createListTable(); err != nil {
		fmt.Println(err)
		fmt.Println("Can not create create list table.")
		return err
	}

	return nil
}

func (d *DBType) createListTable() error {
	query := `
		create table if not exists list (
			id serial primary key,
			createdAt timestamp,
			description varchar(200),
			updatedAt timestamp,
			user_id int references users(id)
		)
	`

	_, err := d.DB.Exec(query)

	return err
}

func (d *DBType) createUserTable() error {
	query := `
		create table if not exists users (
			id serial primary key,
			createdAt timestamp,
			email varchar(60) unique,
			name varchar(30),
			password TEXT,
			updatedAt timestamp
		)
	`

	_, err := d.DB.Exec(query)

	return err
}
