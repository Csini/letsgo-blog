package utils_db

import (
	"config"

	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection(gormconfig *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.GetDbName()), gormconfig)
	if err != nil {
		log.Error(err)
		return nil, errors.New("Failed to connect to the database")
	} else {
		log.Info("Conneted to DB!!")
		//sqlite
		db.Exec("PRAGMA foreign_keys = ON")
		db.Logger = db.Logger.LogMode(logger.Info)
	}
	return db, nil
}
