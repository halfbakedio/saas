package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start SaaS",
	Long:  `Start the SaaS server.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Starting SaaS server...")
	},
}
