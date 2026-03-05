package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	db := connectDB()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback(ctx)

	file, err := excelize.OpenFile("data/features2.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := file.GetRows("Features")
	if err != nil {
		log.Fatal(err)
	}

	header := rows[0]

	// license types are between column 2 and last column - 1
	licenseTypes := header[2 : len(header)-1]

	licenseTypeIDs := map[string]int{}

	// insert license types
	for _, lt := range licenseTypes {

		var id int

		err := tx.QueryRow(ctx,
			`INSERT INTO license_types(name)
			 VALUES($1)
			 ON CONFLICT(name)
			 DO UPDATE SET name = EXCLUDED.name
			 RETURNING id`,
			lt,
		).Scan(&id)

		if err != nil {
			log.Fatal(err)
		}

		licenseTypeIDs[lt] = id
	}

	for i := 1; i < len(rows); i++ {

		row := rows[i]

		featureName := row[0]
		description := row[1]

		// vehicle type is last column
		vehicleType := strings.ToUpper(strings.TrimSpace(row[len(row)-1]))

		if vehicleType == "" {
			vehicleType = "COMMON"
		}

		var featureID int

		err := tx.QueryRow(ctx,
			`INSERT INTO features(name,description,vehicle_type)
			 VALUES($1,$2,$3)
			 ON CONFLICT(name)
			 DO UPDATE SET
			 description = EXCLUDED.description,
			 vehicle_type = EXCLUDED.vehicle_type
			 RETURNING id`,
			featureName,
			description,
			vehicleType,
		).Scan(&featureID)

		if err != nil {
			log.Fatal(err)
		}

		for j, lt := range licenseTypes {

			enabled := false

			val := strings.TrimSpace(row[j+2])

			if strings.ToUpper(val) == "Y" {
				enabled = true
			}

			_, err := tx.Exec(ctx,
				`INSERT INTO license_type_features
				 (license_type_id,feature_id,enabled)
				 VALUES($1,$2,$3)
				 ON CONFLICT(license_type_id,feature_id)
				 DO UPDATE SET enabled = EXCLUDED.enabled`,
				licenseTypeIDs[lt],
				featureID,
				enabled,
			)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Excel import completed successfully")
}

func connectDB() *pgxpool.Pool {

	dbURL := os.Getenv("DB_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connected")

	return pool
}
