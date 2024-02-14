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

	//"entity"

	"gorm.io/gorm"
	utils_db "utils/db"
	//"golang.org/x/crypto/bcrypt"
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
	var gormconfig = &gorm.Config{

		PrepareStmt: false,
	}

	db, err := utils_db.OpenConnection(gormconfig)

	if err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	} else {
		sqlDB, err := db.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v\n", err)
			os.Exit(1)
		}
	}

	log.Info("DB connected")

	//	// Migrate the schema
	//	db.AutoMigrate(&entity.User{}, &entity.Blog{}, &entity.Comment{})
	//
	//	// Create
	//	pw_testuser, _ := HashPassword("testuser1")
	//	db.Create(&entity.User{Username: "testuser", Pw: pw_testuser})
	//
	//	pw_admin, _ := HashPassword("admin1")
	//	db.Create(&entity.User{Username: "admin", Pw: pw_admin})
	//
	//	pw_gipszjakab, _ := HashPassword("gipszjakab1")
	//	db.Create(&entity.User{Username: "gipszjakab", Pw: pw_gipszjakab})
	//
	//	db.Create(&entity.Blog{User_ID: "gipszjakab", Content: "xxxxxx"})
	//
	//	log.Info("DB created")

	AuthenticationAPIService := impl.NewAuthenticationAPIService()
	AuthenticationAPIController := openapi.NewAuthenticationAPIController(AuthenticationAPIService)

	CommentAPIService := impl.NewCommentAPIService()
	CommentAPIController := openapi.NewCommentAPIController(CommentAPIService)

	StatisticsAPIService := impl.NewStatisticsAPIService()
	StatisticsAPIController := openapi.NewStatisticsAPIController(StatisticsAPIService)

	router := openapi.NewRouter(AuthenticationAPIController, CommentAPIController, StatisticsAPIController)

	log.Info("Server started")

	log.Fatal(http.ListenAndServe(":8085", router))
}

/*
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
*/
