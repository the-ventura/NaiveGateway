package main

import (
	"naivegateway/internal/api"
	"naivegateway/internal/database"
	"naivegateway/internal/frontend"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "naivegateway",
	Short: "A payment gateway",
}

func init() {
	rootCmd.AddCommand(api.NewCommand())
	rootCmd.AddCommand(database.NewCommand())
	rootCmd.AddCommand(frontend.NewCommand())
}

func main() {
	rootCmd.Execute()
}
