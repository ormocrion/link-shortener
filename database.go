package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

type Bookmark struct {
	Alias string
	Link string
}

func dbConnect() {
	database, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Невозможно соединиться с базой: %v\n", err)
		os.Exit(1)
	}
	db = database
}

func dbClose() {
	db.Close()
}

func insertBookmark(alias string, link string) {
	_, err := db.Exec("INSERT INTO bookmarks (alias, link) VALUES ($1, $2)", alias, link);
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка сохранения закладки: %v\n", err)
		os.Exit(1)
	}
}

func getBookmark(alias string) string {
	var url string
	
	err := db.QueryRow("SELECT link FROM bookmarks WHERE alias=$1", alias).Scan(&url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения закладки: %v\n", err)
		os.Exit(1)
	}
	
	return url
}

func getAllBookmarks() []byte {
	rows, err := db.Query("SELECT alias, link FROM bookmarks")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения всех закладок: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	bookmarks := []Bookmark{}

	for rows.Next() {
		b := Bookmark{}
		err := rows.Scan(&b.Alias, &b.Link)
		if err != nil {
			fmt.Printf("Ошибка сохранения одной закладки в списке: %v\n", err)
			continue
		}
		bookmarks = append(bookmarks, b)
	}

	json, err := json.Marshal(bookmarks)
	if err != nil {
		fmt.Printf("Ошибка при маршализации закладок: %v\n", err)
	}
	fileSaver(json, err)
	return json
}

