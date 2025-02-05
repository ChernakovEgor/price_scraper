package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ChernakovEgor/price_scraper/internal/database"
	"github.com/ChernakovEgor/price_scraper/internal/scraper"
	_ "github.com/mattn/go-sqlite3"
)

const url1 = "https://chaster-club.ru/catalog/product/orient-ra-ac0p04y"
const url2 = "https://www.ozon.ru/product/roland-fp-30x-wh-tsifrovoe-pianino-beloe-1280857545/?avtc=1&avte=4&avts=1738784206"
const url3 = "https://www.dns-shop.ru/product/9c399838380bed20/videokarta-asus-geforce-rtx-4070-dual-white-oc-edition-dual-rtx4070-o12g-white/"

func main() {
	conn, err := sql.Open("sqlite3", "file:price_scraper.db")
	if err != nil {
		log.Fatalf("connecting to db: %v", err)
	}

	db := database.New(conn)
	scraper := scraper.NewScraper(db)

	res, err := scraper.FetchURL(3, url3)
	fmt.Println(res, err)

	// err = scraper.ProcessURLs()
	// if err != nil {
	// 	log.Printf("error processing URLS: %v", err)
	// 	os.Exit(1)
	// }

	os.Exit(0)
}
