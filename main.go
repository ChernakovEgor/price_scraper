package main

import (
	"database/sql"
	"log"

	"github.com/ChernakovEgor/price_scraper/internal/database"
	"github.com/ChernakovEgor/price_scraper/internal/scraper"
	_ "github.com/mattn/go-sqlite3"
)

const url = "https://chaster-club.ru/catalog/product/orient-ra-ac0p04y"

func main() {
	conn, err := sql.Open("sqlite3", "file:price_scraper.db")
	if err != nil {
		log.Fatalf("connecting to db: %v", err)
	}

	db := database.New(conn)
	res, err := scraper.FetchURL(db, 1, url)
	if err != nil {
		log.Fatalf("scraping url '%s': %v", url, err)
	}

	log.Printf("inserted at row %d", res)
}
