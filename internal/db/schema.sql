-- schema.sql

-- Components schema
-- DROP TABLE IF EXISTS component CASCADE;
CREATE TABLE IF NOT EXISTS component (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INTEGER REFERENCES component(id) ON DELETE SET NULL -- Allows NULL parent_id for root components, SET NULL on parent delete
);