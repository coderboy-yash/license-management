package repository

import (
	"context"
	"license-service/config"
	"time"

	"github.com/jackc/pgx/v5"
)

type LicenseRepository struct{}

func NewLicenseRepository() *LicenseRepository {
	return &LicenseRepository{}
}

func (r *LicenseRepository) GetLicenseTypeID(ctx context.Context, name string) (int, error) {

	var id int

	err := config.DB.QueryRow(ctx,
		`SELECT id FROM license_types WHERE name=$1`,
		name,
	).Scan(&id)

	return id, err
}

func (r *LicenseRepository) InsertLicense(
	ctx context.Context,
	licenseID string,
	vin string,
	vehicleType string,
	licenseTypeID int,
	expiry time.Time,
) error {

	_, err := config.DB.Exec(ctx,
		`INSERT INTO licenses
		(license_id, vin, vehicle_type, license_type_id, expiry_date)
		VALUES ($1,$2,$3,$4,$5)`,
		licenseID,
		vin,
		vehicleType,
		licenseTypeID,
		expiry,
	)

	return err
}

// ////
func (r *LicenseRepository) GetLicenseFeatures(ctx context.Context, licenseID string) (pgx.Rows, error) {

	query := `
	SELECT 
    f.name,
    f.description,
    ltf.enabled
FROM licenses l
JOIN license_type_features ltf
    ON l.license_type_id = ltf.license_type_id
JOIN features f
    ON f.id = ltf.feature_id
WHERE l.license_id = $1
AND l.expiry_date > NOW()
AND (
    f.vehicle_type = 'COMMON'
    OR f.vehicle_type = l.vehicle_type
);
	`

	return config.DB.Query(ctx, query, licenseID)
}
