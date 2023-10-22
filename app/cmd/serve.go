package cmd

import (
	"github.com/fhlarif/goblog/app/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
	Long:  "Application will be served on the given host and port from the config.yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	bootstrap.InitBootstrap()
}
