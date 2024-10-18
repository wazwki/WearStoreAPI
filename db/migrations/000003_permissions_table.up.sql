CREATE TABLE permissions_table (
    email INT REFERENCES users_table(email),
    role VARCHAR(50) NOT NULL
);