/*
 * Let's go Blog API
 *
 * Application providing Blog.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package impl

import (
	"context"
	"errors"
	//"net/http"

	log "github.com/sirupsen/logrus"

	openapi "generated/openapi"

	"config"
	"entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	utils_jwt "utils/jwt"

	"golang.org/x/crypto/bcrypt"
	//"time"
)

// AuthenticationAPIService is a service that implements the logic for the AuthenticationAPIServicer
// This service should implement the business logic for every endpoint for the AuthenticationAPI API.
// Include any external packages or services that will be required by this service.
type AuthenticationAPIService struct {
}

// NewAuthenticationAPIService creates a default api service
func NewAuthenticationAPIService() openapi.AuthenticationAPIServicer {
	return &AuthenticationAPIService{}
}

// PostLogin - provides JWT token
func (s *AuthenticationAPIService) PostLogin(ctx context.Context, loginRequest openapi.LoginRequest) (openapi.ImplResponse, error) {
	// TODO - update PostLogin with the required logic for this service method.
	// Add api_authentication_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, string{}) or use other options such as http.Ok ...
	// return Response(200, string{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	db, err := gorm.Open(sqlite.Open(config.GetDbName()), &gorm.Config{

		PrepareStmt: false,
	})
	if err != nil {
		log.Error(err)
		return openapi.Response(500, nil), errors.New("PostLogin failed to connect database")
	} else {
		log.Info("yuhuuuuu")
		//sqlite
		db.Exec("PRAGMA foreign_keys = ON")
	}

	var user entity.User
	// Get
	result := db.Table("users").Select("username, pw").Where("username=?", loginRequest.User).Scan(&user)
	//result.RowsAffected // returns found records count, equals `len(users)`
	//result.Error        // returns error
	if result.Error != nil {
		return openapi.Response(401, nil), errors.New("PostLogin - bad User or pw")
	}

	if checkPasswordHash(loginRequest.Pw, user.Pw) {

		token, err := utils_jwt.GenerateJWT(user.Username)

		if err != nil {
			log.Error(err)
			return openapi.Response(401, nil), errors.New("PostLogin - Token generation error")
		}

		return openapi.Response(200, token), nil
	}

	return openapi.Response(401, nil), errors.New("PostLogin - bad User or pw")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
