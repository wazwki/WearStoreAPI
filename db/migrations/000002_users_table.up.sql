CREATE TABLE IF NOT EXISTS users_table(
    id SERIAL PRIMARY KEY,
    email VARCHAR(150) UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL
);
