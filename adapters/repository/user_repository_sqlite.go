package repository

import (
	"database/sql"
	"log"

	"github.com/YoshikiGit/poc-go-clean-architecture/entities"
	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func NewUserSqlite(db string) UserSqlite {
	return UserSqlite{
		db: db,
	}
}

type UserSqlite struct {
	db string
}

func (a UserSqlite) Create(user entities.User) (entities.User, error) {
	DbConnection, _ := sql.Open("sqlite3", a.db)
	defer DbConnection.Close()

	cmd := "INSERT INTO user (id, name, age, created_at) VALUES (?, ?, ?, ?)"
	_, err := DbConnection.Exec(cmd, user.ID(), user.Name(), user.Age(), user.CreatedAt())
	if err != nil {
		log.Fatalln(err)
	}
	return user, nil
}
