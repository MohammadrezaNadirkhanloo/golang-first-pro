package cmd

import (
	"first-app/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	configs := ConfigsSet()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"app":     viper.Get("App.Name"),
		})
	})

	if err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func ConfigsSet() config.Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error read config file")
	}

	var configs config.Config

	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	return configs
}
