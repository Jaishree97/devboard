CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO projects (name)
SELECT 'DevOps Learning'
WHERE NOT EXISTS (
    SELECT 1
    FROM projects
    WHERE name = 'DevOps Learning'
);