package main

import (
	"context"
	"fmt"
	"log"

	"github.com/halfbakedio/saas/config"
	"github.com/halfbakedio/saas/ent"

	_ "github.com/lib/pq"
)

func main() {
	c := config.GetConfig()
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Db.Host,
		c.Db.Port,
		c.Db.User,
		c.Db.Name,
		c.Db.Password,
		c.Db.SslMode,
	)

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	log.Print("Connected to Postgres")

	defer client.Close()

	// run auto migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Print("Schema created successfully")
}
