gooseUp:
	cd ./sql/schema && goose sqlite3 ../../price_scraper.db up

gooseDown:
	cd ./sql/schema && goose sqlite3 ../../price_scraper.db down

sqlc:
	cd ./sql/queries/ && sqlc generate

test:
	go test ./...
