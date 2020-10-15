package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "auth-server",
		Short: "auth-server",
		Long:  "Simple and Powerful Golang Auth Server",
		Run:   handleExecute,
	}
)

func handleExecute(cmd *cobra.Command, args []string) {
	fmt.Println("Test")
}

func Execute() error {
	return rootCmd.Execute()
}
