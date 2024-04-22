CREATE TABLE IF NOT EXISTS sessions (
    id varchar not null,
    token varchar not null,
    is_active bool default false,
    user_id varchar not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

