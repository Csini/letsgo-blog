/*
 * Let's go Blog API
 *
 * Application providing Blog.
 *
 * API version: 1.0.0
 */

package main

import (
	openapi "generated/openapi"
	impl "impl"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the info severity or above.
	//log.SetLevel(log.InfoLevel)
	//TODO
	log.SetLevel(log.DebugLevel)
}

func main() {

	/*log.WithFields(log.Fields{
		 "prefix":      "sensor",
		 "temperature": -4,
	 }).Info("Temperature changes")*/
	var err error
	var db *gorm.DB
	if db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	} else {
		sqlDB, err := db.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v\n", err)
		}

		//sqlite
		db.Exec("PRAGMA foreign_keys = ON")

		db.Logger = db.Logger.LogMode(logger.Info)
	}

	// Migrate the schema
	db.AutoMigrate(&entity.User{}, &entity.Blog{}, &entity.Comment{})

	// Create
	db.Create(&entity.User{Username: "test", Pw: "abc"})

	log.Info("DB created")

	StatisticsAPIService := impl.NewStatisticsAPIService()
	StatisticsAPIController := openapi.NewStatisticsAPIController(StatisticsAPIService)

	router := openapi.NewRouter(StatisticsAPIController)

	log.Info("Server started")

	log.Fatal(http.ListenAndServe(":8085", router))
}
