package main

import (
	"github.com/spf13/cobra"
)

const rootCmd = &cobra.Command{
	Use:   "auth-server",
	Short: "auth-server",
	Long:  `Simple and Powerful Golang Auth Server`,
	Run: handleExecute
}

func handleExecute(cmd *cobra.Command, args []string) {
	// do stuff anything
}

func Execute() {
	if err := rootCmd.Execute
}
