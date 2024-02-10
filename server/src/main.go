package main

import (
	"common"
	log "github.com/sirupsen/logrus"
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

	common.PrintHello("universe")

	log.WithFields(log.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")

}
