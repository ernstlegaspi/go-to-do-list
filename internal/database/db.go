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
			updatedAt timestamp
		)
	`

	_, err := d.DB.Exec(query)

	return err
}
