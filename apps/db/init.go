package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

func DbInit() *gorm.DB {
	log.Info("Starting database...")

	dbURI := "host=localhost user=postgres password=admin dbname=go_todo port=5432"

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Panic("Failed to open database" + err.Error())
	}

	db.LogMode(true)
	return db

}