package cmd

import (
	"github.com/halfbakedio/saas/config"
	"github.com/halfbakedio/saas/constants"
	"github.com/halfbakedio/saas/db"
	"github.com/halfbakedio/saas/service"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	destroy     bool
	userService *service.UserService
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database",
	Long:  "Seed the SaaS server.",
	Run:   seed,
}

func init() {
	seedCmd.PersistentFlags().BoolVarP(
		&destroy,
		"destroy",
		"d",
		false,
		"Destroy all data in the database before seeding",
	)

	_ = viper.BindPFlag("destroy", seedCmd.PersistentFlags().Lookup("destroy"))

	viper.SetDefault("destroy", false)
}

func seed(cmd *cobra.Command, args []string) {
	cfg := config.GetConfig()

	if !constants.IsDevelopment(cfg.Env) {
		log.Fatal("Seeding is only allowed in development environment.")
	}

	log.Debug("Seeding SaaS server...")

	conn := db.GetConnection()

	err := conn.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	userService = service.NewUserService(conn)

	seedUsers()
}

func seedUsers() {
	users := []map[string]string{
		{"email": "admin@saas.io", "password": "password"},
		{"email": "user@saas.io", "password": "password"},
	}

	for _, user := range users {
		if destroy {
			log.Debugf("Destroying user %s", user["email"])
			err := userService.DeleteUserByEmail(user["email"])
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err := userService.FindOrCreateUser(
			user["email"],
			map[string]interface{}{"password": user["password"]},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
