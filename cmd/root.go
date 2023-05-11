package cmd

import (
	"log"
	"os"

	"github.com/deepfake-db/internal/config"
	eventListener "github.com/deepfake-db/internal/eventListener"
	"github.com/deepfake-db/internal/routes"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "deepfake-db",
	Version: version,
	Short:   "Starts the db",
	Args:    cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		config.LoadClient()
		config.InitialMigrationDB()
	},
	Run: func(cmd *cobra.Command, args []string) {
		r := routes.ConfigRouter()
		go routes.RunRouter(r)
		eventListener.Listen()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("error starting db: ", err)
		os.Exit(1)
	}
}
