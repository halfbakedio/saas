package main

import (
	"github.com/halfbakedio/saas/cmd"
	"github.com/halfbakedio/saas/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	initLogging()

	cmd.Execute()
}

func initLogging() {
	cfg := config.GetConfig()

	if logLevel, err := log.ParseLevel(cfg.Log.Level); err == nil {
		log.SetLevel(logLevel)
	}

	if cfg.Log.Formatter == "json" {
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}
