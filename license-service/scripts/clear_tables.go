package scripts

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	db := connectDB()
	defer db.Close()

	ctx := context.Background()

	_, err = db.Exec(ctx, `
	TRUNCATE TABLE
	license_type_features,
	licenses,
	license_types,
	features
	RESTART IDENTITY
	CASCADE
	`)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("All tables cleared successfully")
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
