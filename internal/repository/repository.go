package repository

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Migration для проверки созданной струтуры
func Migration(dbconnection *sql.DB) {
	query, err := ioutil.ReadFile("internal/repository/tables.sql")
	if err != nil {
		log.Fatal(err)
	}
	s, err := dbconnection.Prepare(string(query))
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.Exec()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

// Connect Возвращает сооединение с БД
func Connect(dbName string) *sql.DB {
	database, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	return database
}
