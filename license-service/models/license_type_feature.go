package models

type LicenseTypeFeature struct {
	ID            int  `json:"id"`
	LicenseTypeID int  `json:"license_type_id"`
	FeatureID     int  `json:"feature_id"`
	Enabled       bool `json:"enabled"`
}
