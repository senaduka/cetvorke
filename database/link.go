package database

import (
	"fmt"
)

type Link struct {
	ID               int
	TitleSlug        string `db:"title_slug""`
	Year, Month, Day int
	Title            string `db:"post_title"`
}

func (l *Link) GemtextLink() string {
	return fmt.Sprintf("=> /clanak/%d/%s %d-%d-%d - %s\n", l.ID, l.TitleSlug, l.Year, l.Month, l.Day, l.Title)
}

func GetRecentLinks() ([]Link, error) {
	db := GetDB()
	links := []Link{}
	err := db.Select(&links, "SELECT id, title_slug, year, month, day, post_title FROM posts ORDER BY date DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	return links, nil
}
