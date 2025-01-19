package cmd

import (
	"github.com/0125nia/Mercury/ipconf"
	"github.com/spf13/cobra"
)

var ipconfCommand = &cobra.Command{
	Use:   "ipconf",
	Short: "Mercury ipconf",
	Run: func(cmd *cobra.Command, args []string) {
		ipconf.RunMain()
	},
}

func init() {
	rootCmd.AddCommand(ipconfCommand)
}
