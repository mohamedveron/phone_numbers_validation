package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


type Persistence struct {
	sourcePath string
	database *sql.DB
}

func NewPersistence(sourcePath string) *Persistence{
	database, err := sql.Open("sqlite3", sourcePath)

	if err != nil {
		fmt.Print(err)
	}

	return &Persistence{
		sourcePath: sourcePath,
		database: database,
	}
}



