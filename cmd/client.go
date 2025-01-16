package cmd

import (
	"github.com/0125nia/Mercury/client"
	"github.com/spf13/cobra"
)

var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "Mercury client",
	Run: func(cmd *cobra.Command, args []string) {
		client.RunMain()
	},
}

func init() {
	rootCmd.AddCommand(clientCommand)
}
