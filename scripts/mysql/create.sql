DROP TABLE IF EXISTS user;

CREATE TABLE IF NOT EXISTS user (
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);