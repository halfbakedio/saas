package main

import (
	"context"
	"log"

	"github.com/halfbakedio/saas/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=saas password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
