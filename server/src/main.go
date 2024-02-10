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
	log "github.com/sirupsen/logrus"
	impl "impl"
	"net/http"
	"os"
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

	log.Info("Server started")

	StatisticsAPIService := impl.NewStatisticsAPIService()
	StatisticsAPIController := openapi.NewStatisticsAPIController(StatisticsAPIService)

	router := openapi.NewRouter(StatisticsAPIController)
	log.Fatal(http.ListenAndServe(":8085", router))
}
