package main

import (
	"naivegateway/internal/api"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "naivegateway",
	Short: "A payment gateway",
}

func init() {
	rootCmd.AddCommand(api.NewCommand())
}

func main() {
	rootCmd.Execute()
}
