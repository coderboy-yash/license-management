package service

import (
	"context"
	"time"

	"license-service/models"
	"license-service/repository"

	"github.com/google/uuid"
)

type LicenseService struct {
	repo *repository.LicenseRepository
}

func NewLicenseService(r *repository.LicenseRepository) *LicenseService {
	return &LicenseService{repo: r}
}

func (s *LicenseService) CreateLicense(ctx context.Context, req models.CreateLicenseRequest) (string, error) {

	licenseTypeID, err := s.repo.GetLicenseTypeID(ctx, req.LicenseType)
	if err != nil {
		return "", err
	}

	expiry, err := time.Parse("2006-01-02", req.ExpiryDate)
	if err != nil {
		return "", err
	}

	licenseID := uuid.New().String()

	err = s.repo.InsertLicense(ctx, licenseID, req.VIN, req.VehicleType, licenseTypeID, expiry)
	if err != nil {
		return "", err
	}

	return licenseID, nil
}

// ///
func (s *LicenseService) GetLicenseFeatures(ctx context.Context, licenseID string) ([]models.FeatureResponse, error) {

	rows, err := s.repo.GetLicenseFeatures(ctx, licenseID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var features []models.FeatureResponse

	for rows.Next() {

		var f models.FeatureResponse

		err := rows.Scan(&f.Name, &f.Description, &f.Enabled)
		if err != nil {
			return nil, err
		}

		features = append(features, f)
	}

	return features, nil
}
