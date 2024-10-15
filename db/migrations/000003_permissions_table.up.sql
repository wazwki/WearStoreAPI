CREATE TABLE permissions_table (
    user_id INT REFERENCES users(id),
    role VARCHAR(50) NOT NULL
);