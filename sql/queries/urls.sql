-- name: AddURL :one
INSERT INTO urls(url)
     VALUES (?)
  RETURNING *;

-- name: GetURLs :many
SELECT id, url
  FROM urls;
