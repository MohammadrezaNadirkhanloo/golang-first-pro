package cmd

import (
	"first-app/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "Server app on dev server",
	Long:  `Application will be served on host and port defined in config.yml file`,
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func server() {
	bootstrap.Server()
}
