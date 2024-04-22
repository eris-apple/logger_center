CREATE TABLE IF NOT EXISTS users (
    id varchar not null,
    email varchar not null unique,
    password varchar not null,
    role varchar not null default 'guest',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);