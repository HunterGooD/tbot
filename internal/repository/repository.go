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

type User struct {
	ID       int
	Username string
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

// Connect сооединение с БД
func (r *Repo) Connect() {
	database, err := sql.Open("sqlite3", r.dbName)
	if err != nil {
		log.Fatal(err)
	}
	r.dbconnect = database
}

func (r *Repo) GetIDUsers() []int {
	var result = make([]int, 0)
	rows, err := r.dbconnect.Query("SELECT id FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var a int
		if err = rows.Scan(&a); err != nil {
			log.Fatal(err)
		}
		result = append(result, a)
	}
	return result
}

// GetUserByID ..
func (r *Repo) GetUserByID(id int) *User {
	// var (
	// 	idDB     int
	// 	username string
	// )
	rows, err := r.dbconnect.Query("SELECT * FROM user WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		var u = User{}
		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			log.Fatal(err)
		}
		return &u
	}
	return nil
}

// AddUser ..
func (r *Repo) AddUser(id int, username string) error {
	_, err := r.dbconnect.Exec("INSERT INTO USER (id, username) VALUES ($1, $2)", id, username)
	if err != nil {
		return err
	}
	return nil
}
