package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"mytest.net/restapi/config"
	"mytest.net/restapi/api/model"
	"mytest.net/restapi/api/db"
	"mytest.net/restapi/api/router"
	"runtime/debug"
	"os"
)

var appConfig config.Config

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	err := envconfig.Process("restapi", &appConfig)
	if err != nil {
		log.Error("Failed to populate config from environment variables")
	}
}

func main() {

	defer func() { 
		if r := recover(); r != nil {
			log.WithField("stacktrace", string(debug.Stack())).Errorf("failed due to panic. %+v", r)
		}
	}()
	cl := new(model.Conf)
	c := new(model.Clients)
	// initialize application configurations
	err := envconfig.Process("app", &appConfig)
	if err != nil {
		log.Error("Failed to populate application configurations from environment variables")
		return
	}

	c.Config = appConfig
	cl.Config = appConfig
	// initialize db connection
	dbClient, err := db.Connect(appConfig.Database, appConfig.DatabaseURL)
	if err != nil {
		log.Errorf("Failed to connect to mongo client, error: %v", err)
		return
	}
	c.DB = dbClient

	r := router.SetupRouter(c)
	r.Run(":4000")
	log.Println("Starting server on :4000")
}