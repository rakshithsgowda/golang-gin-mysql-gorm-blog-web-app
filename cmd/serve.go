package cmd

import (
	"blog/pkg/config"
	"blog/pkg/routing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	
	config.Set()


	// initialize roter
	routing.Init()

	router:= routing.GetRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})
	routing.Serve()
}

