package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func GetDB() *sqlx.DB {
	once.Do(func() {
		db = sqlx.MustConnect("sqlite3", "./cetvorke.db")
		postsTable := `
    CREATE TABLE IF NOT EXISTS posts(
        id INTEGER NOT NULL PRIMARY KEY,
        date DATETIME,
        year INTEGER,
        month INTEGER,
        day INTEGER,
        post_title TEXT,
        title_slug TEXT,
        markdown_content TEXT,
        gemtext_content TEXT
     );
    `
		db.Exec(postsTable)
	})

	return db
}
