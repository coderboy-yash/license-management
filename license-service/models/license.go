package models

import "time"

type CreateLicenseRequest struct {
	VIN         string `json:"vin"`
	LicenseType string `json:"license_type"`
	VehicleType string `json:"vehicle_type"`
	ExpiryDate  string `json:"expiry_date"`
}

type License struct {
	LicenseID     string
	VIN           string
	VehicleType   string
	LicenseTypeID int
	ExpiryDate    time.Time
}
