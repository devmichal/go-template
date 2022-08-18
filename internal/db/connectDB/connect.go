package connectDB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	postgresDsnScheleton string = "host=%s user=%s password=%s dbname=%s port=%s"
)

func Connect() (db *gorm.DB) {

	dns := fmt.Sprintf(
		postgresDsnScheleton,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
