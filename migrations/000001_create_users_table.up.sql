-- +migrate Up
    CREATE TABLE users
    (
        id         INT AUTO_INCREMENT PRIMARY KEY,
        name       VARCHAR(255),
        email      VARCHAR(255) UNIQUE,
        password   VARCHAR(255),
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );