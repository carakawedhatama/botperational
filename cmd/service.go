package cmd

import (
	"fmt"
	"os"

	"botperational/config"
	"botperational/internal/bootstrap"

	"github.com/joho/godotenv"
	"github.com/runsystemid/golog"
	"github.com/spf13/cobra"
)

var (
	configFile string
	command    = &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Run service",
		Run: func(c *cobra.Command, args []string) {
			// Load env variable
			err := godotenv.Load(".env")
			if err != nil {
				fmt.Println("Fatal error loading .env file.\n", err)
				os.Exit(1)
			}

			// Load configuration
			conf := config.Config{}
			conf.LoadConfig("config")

			// Initialize Logger
			loggerConfig := golog.Config{
				App:             conf.App,
				AppVer:          conf.AppVer,
				Env:             conf.Env,
				FileLocation:    conf.Log.FileLocation,
				FileTDRLocation: conf.Log.FileTDRLocation,
				FileMaxSize:     conf.Log.FileMaxSize,
				FileMaxBackup:   conf.Log.FileMaxBackup,
				FileMaxAge:      conf.Log.FileMaxAge,
				Stdout:          conf.Log.Stdout,
			}
			golog.Load(loggerConfig)

			bootstrap.Run(&conf)
		},
	}
)

func GetCommand() *cobra.Command {
	command.Flags().StringVar(&configFile, "config", "./config.yaml", "Set config file path")

	return command
}
