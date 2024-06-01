package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/halfbakedio/saas/config"
	"github.com/halfbakedio/saas/ent"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Connection struct {
	Ctx    context.Context
	Client *ent.Client
}

var (
	lock     = &sync.Mutex{}
	instance *Connection
)

func GetConnection() *Connection {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Connection{
				Ctx: context.Background(),
			}
		}
	}

	log.Tracef("config: %+v", instance)

	return instance
}

func (c *Connection) Open() error {
	if c.Client != nil {
		return nil
	}

	cfg := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password,
		cfg.DB.SslMode,
	)

	log.Debug("Connected to Postgres")

	drv, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	c.Client = ent.NewClient(ent.Driver(drv))

	return nil
}

func (c *Connection) Close() {
	if c.Client == nil {
		return
	}

	c.Client.Close()
}

func (c *Connection) AutoMigrate() {
	if err := c.Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Info("Schema created successfully")
}
