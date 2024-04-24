CREATE TABLE IF NOT EXISTS service_accounts (
    id varchar not null,
    project_id varchar null,
    is_active bool default false,
    secret varchar null,
    name varchar null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

