CREATE TABLE license_types (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
ALTER TABLE license_types ADD CONSTRAINT license_types_name_unique UNIQUE(name);