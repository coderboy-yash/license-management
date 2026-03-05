CREATE TABLE features (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    vehicle_type TEXT DEFAULT 'COMMON',
    created_at TIMESTAMP DEFAULT NOW()
);
ALTER TABLE features ADD CONSTRAINT features_name_unique UNIQUE(name);