package cmd

import (
	"github.com/0125nia/Mercury/perf"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(perfCmd)
	perfCmd.PersistentFlags().Int32Var(&perf.TcpConnNum, "tcp_conn_num", 10000, "tcp connection number")
}

var perfCmd = &cobra.Command{
	Use: "perf",
	Run: func(cmd *cobra.Command, args []string) {
		perf.RunMain()
	},
}
