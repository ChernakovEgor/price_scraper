-- name: AddURL :one
INSERT INTO urls(url, tracked_element)
     VALUES (?, ?)
  RETURNING *;
