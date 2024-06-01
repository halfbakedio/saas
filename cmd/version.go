package cmd

import (
	"fmt"

	"github.com/halfbakedio/saas/pkg"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of saas",
	Long:  `SaaS version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.VERSION)
	},
}
