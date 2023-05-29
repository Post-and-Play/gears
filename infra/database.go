package infra

import (
	"fmt"
	"log"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/Post-and-Play/gears/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DevDatabaseConnect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	ssl := os.Getenv("SSL")

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, dbport, ssl)
	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Local database connection error")
	}
	autoMigrateModels()
}

func ProdDatabaseConnect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	ssl := os.Getenv("SSL")

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, dbport, ssl)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        conn,
	}))

	if err != nil {
		log.Panic("Prod database connection error")
	}
	autoMigrateModels()
}

func autoMigrateModels() {
	DB.AutoMigrate(&models.User{}, &models.Game{}, &models.Review{}, &models.Like{})
}
