package database

import (
	"database/sql"
	"fmt"
)

type Post struct {
	ID               int
	Title            string `db:"post_title"`
	Content          string
	Date             string
	MarkdownContent  string         `db:"markdown_content"`
	GemtextContent   sql.NullString `db:"gemtext_content"`
	TitleSlug        string         `db:"title_slug"`
	Year, Month, Day int
}

func (p *Post) GemtextPage() string {
	return fmt.Sprintf("# %s\n\n", p.Title) + p.GemtextContent.String
}

func GetPost(id int) (*Post, error) {
	db := GetDB()
	post := &Post{}
	err := db.Get(post, "SELECT id, post_title, date, markdown_content, gemtext_content, title_slug, year, month, day FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
