CREATE TABLE IF NOT EXISTS project (
    id serial PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    description TEXT DEFAULT NULL
);

