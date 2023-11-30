package cmd

import (
	"blog/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
  Use:   "serve",
  Short: "Serve app on dev server",
  Long:  `Application will be served on the host and port defined in config.yml in the config file`,
  Run: func(cmd *cobra.Command, args []string) {
    serve()
  },
}


func serve() {
	bootstrap.Serve()
	
}

