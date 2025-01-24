-- +goose Up
CREATE TABLE parsing_results (
  id INTEGER PRIMARY KEY NOT NULL,
  url_id INTEGER REFERENCES ulrs(id) NOT NULL,
  date_run DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  status_code INT NOT NULL,
  raw_body TEXT,
  target_field NUMERIC
);

-- +goose Down
DROP TABLE parsing_results;
