// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: urls.sql

package database

import (
	"context"
)

const addURL = `-- name: AddURL :one
INSERT INTO urls(url)
     VALUES (?)
  RETURNING id, url, created_at, updated_at
`

func (q *Queries) AddURL(ctx context.Context, url string) (Url, error) {
	row := q.db.QueryRowContext(ctx, addURL, url)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getURLs = `-- name: GetURLs :many
SELECT id, url
  FROM urls
`

type GetURLsRow struct {
	ID  int64
	Url string
}

func (q *Queries) GetURLs(ctx context.Context) ([]GetURLsRow, error) {
	rows, err := q.db.QueryContext(ctx, getURLs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetURLsRow
	for rows.Next() {
		var i GetURLsRow
		if err := rows.Scan(&i.ID, &i.Url); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
