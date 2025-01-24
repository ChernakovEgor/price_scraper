-- name: AddParsingResult :one
INSERT INTO parsing_results(url_id, status_code, raw_body)
     VALUES (?, ?, ?)
  RETURNING *;

-- name: GetURLid :one
SELECT id
  FROM urls
 WHERE url = ?;
