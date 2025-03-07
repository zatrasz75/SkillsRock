
-- +migrate Up
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT DEFAULT 'new' CHECK (status IN ('new', 'in_progress', 'done')),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS tasks;