package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "net-sentinel",
	Short: "Net Sentinel CLI an network traffic monitoring tool",
	Long: "Net Sentinel CLI an network traffic monitoring tool",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
