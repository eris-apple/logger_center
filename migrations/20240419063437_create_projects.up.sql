CREATE TABLE IF NOT EXISTS projects (
    id varchar not null,
    name varchar not null,
    prefix varchar not null,
    is_active bool default false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);