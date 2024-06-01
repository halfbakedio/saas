package main

import (
	"github.com/halfbakedio/saas/cmd"
	"github.com/halfbakedio/saas/db"

	log "github.com/sirupsen/logrus"
)

func main() {
	conn := db.GetConnection()

	err := conn.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	cmd.Execute()
}
