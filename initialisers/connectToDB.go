package initialisers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// connect to local database on port 5432
	// dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}