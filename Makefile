gooseUp:
	cd ./sql/schema && goose sqlite3 ../../price_scraper.db up

gooseSeed:
	cd ./sql/seed && goose sqlite3 ../../price_scraper.db up
gooseUnseed:
	cd ./sql/seed && goose sqlite3 ../../price_scraper.db down

gooseDown:
	cd ./sql/schema && goose sqlite3 ../../price_scraper.db down

sqlc:
	cd ./sql/queries/ && sqlc generate

test:
	go test ./...

run:
	go run .
