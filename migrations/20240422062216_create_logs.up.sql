CREATE TABLE IF NOT EXISTS logs (
    id varchar not null,
    chain_id varchar null,
    project_id varchar null,
    content varchar null,
    level varchar null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

