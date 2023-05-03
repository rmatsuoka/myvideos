package model

import (
	"database/sql"
	"io/fs"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//go:generate stringer -type Attr

type Tag struct {
	ID   string
	Name string
	N    int
}

type Model struct {
	db *sql.DB
	fs fs.FS
}

func (m *Model) FS() fs.FS {
	return m.fs
}

func NewModel() *Model {
	db, err := sql.Open("sqlite3", "./myvideos.db")
	if err != nil {
		panic(err)
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return &Model{db, os.DirFS(homedir)}
}
