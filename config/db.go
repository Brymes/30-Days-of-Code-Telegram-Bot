package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBClient *gorm.DB
)

func InitDb() {
	var err error
	uri1 := os.Getenv("DATABASE_URI")

	if uri1 == "" {
		uri1 = "root:mypass@tcp(127.0.0.1:3306)/rewards_cwa?charset=utf8mb4&parseTime=True&loc=Local"
		//uri = "root:mypass@tcp(host.docker.internal:3306)/rewards_cwa?charset=utf8mb4&parseTime=True&loc=Local"
	}

	DBClient, err = gorm.Open(mysql.Open(uri1), &gorm.Config{})
	if err != nil {
		log.Println("Error Connecting to Database")
		log.Fatal(err)
	}
}
