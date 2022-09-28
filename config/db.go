package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBClient *gorm.DB
)

func InitDb() {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=enable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

	host, user, password, dbname, port := "127.0.0.1", "postgres", "mysecretpassword", "postgresDB", 5456
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		host, user, password, dbname, port,
	)

	DBClient, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Println("Error Connecting to Database")
		log.Fatal(err)
	}
	return
}
