package cmd

import (
	"log"

	"github.com/chiragsoni81245/net-sentinel/internal/config"
	"github.com/chiragsoni81245/net-sentinel/internal/server"
	"github.com/spf13/cobra"
)

func runServer(configPath string) {
    // Generate application configuration
    config, err := config.GetConfig(configPath)
    if err != nil {
        log.Fatal(err)
    }

    err = server.NewServer(config)
    if err != nil {
        log.Fatal(err)
    }
}

var configPath string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the net sentinel service",
	Run: func(cmd *cobra.Command, args []string) {
        runServer(configPath)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&configPath, "config", "", "Path to the config file (required)")
	startCmd.MarkFlagRequired("config")
}

