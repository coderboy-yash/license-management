CREATE TABLE licenses (
    license_id UUID PRIMARY KEY,
    vin VARCHAR(20) UNIQUE NOT NULL,
    license_type_id INT REFERENCES license_types(id),
    expiry_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);