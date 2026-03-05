CREATE TABLE license_type_features (
    id SERIAL PRIMARY KEY,
    license_type_id INT REFERENCES license_types(id) ON DELETE CASCADE,
    feature_id INT REFERENCES features(id) ON DELETE CASCADE,
    enabled BOOLEAN DEFAULT FALSE,
    UNIQUE(license_type_id, feature_id)
);

ALTER TABLE license_type_features
ADD CONSTRAINT license_type_feature_unique
UNIQUE(license_type_id, feature_id);