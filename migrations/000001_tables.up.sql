CREATE TABLE IF NOT EXISTS urls(
    id SERIAL PRIMARY KEY,
    hash VARCHAR(10) UNIQUE,
    original_url TEXT
);
