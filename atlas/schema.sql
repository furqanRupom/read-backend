CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
ALTER TABLE authors ADD COLUMN password TEXT;
ALTER TABLE authors ADD COLUMN email TEXT;

