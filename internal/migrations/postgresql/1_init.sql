-- +goose Up
CREATE TABLE projects (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at timestamp default current_timestamp
);

CREATE TABLE goods (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    project_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    priority BIGINT NOT NULL,
    removed BOOL DEFAULT FALSE,
    created_at timestamp default current_timestamp,
    FOREIGN KEY (project_id)
        REFERENCES projects (id)
        ON DELETE CASCADE
);

CREATE INDEX goods_name
ON goods (project_id, name);

INSERT INTO projects (name)
VALUES ('Первая запись');

-- +goose Down
DROP TABLE goods;
DROP TABLE projects;
