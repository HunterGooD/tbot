package repository

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// NewRepo ..
func NewRepo(dbname string) *Repo {
	return &Repo{
		dbName: dbname,
	}
}

// Repo ..
type Repo struct {
	dbconnect *sql.DB
	dbName    string
}

// Migration для проверки созданной струтуры
func (r *Repo) Migration() {
	query, err := ioutil.ReadFile("internal/repository/tables.sql")
	if err != nil {
		log.Fatal(err)
	}
	s, err := r.dbconnect.Prepare(string(query))
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
func (r *Repo) Connect() {
	database, err := sql.Open("sqlite3", r.dbName)
	if err != nil {
		log.Fatal(err)
	}
	r.dbconnect = database
}
