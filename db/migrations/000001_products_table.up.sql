CREATE TABLE IF NOT EXISTS poducts_table(
    id SERIAL PRIMARY KEY,
    price NUMERIC(10, 2) NOT NULL,
    title TEXT NOT NULL,
    photo TEXT,
    description JSONB
);