CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,                         -- auto-incrementing integer ID
    name VARCHAR(255) NOT NULL,                    -- group name
    currency VARCHAR(10) NOT NULL,                 -- e.g., USD, EUR, NGN
    description TEXT,                              -- optional description
    created_by INT NOT NULL,                       -- user ID of creator
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- time of creation
);
