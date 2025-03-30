package config

import (
	"exchangeapp_train/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := Appconfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database, go error: %v", err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(Appconfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Appconfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to configure database, got error: %v", err)
	}
	global.Db = db
}
