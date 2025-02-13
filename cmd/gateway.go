package cmd

import (
	"github.com/0125nia/Mercury/gateway"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Mercury gateway",
	Run: func(cmd *cobra.Command, args []string) {
		gateway.RunMain(ConfigPath)
	},
}
