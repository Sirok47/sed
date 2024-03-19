package dbHandler

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const dsn = "host=localhost user=postgres password=qpwoeirutyM123 dbname=cliDB port=5432 sslmode=disable TimeZone=Europe/London"

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
