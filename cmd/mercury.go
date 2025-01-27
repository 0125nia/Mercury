package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&ConfigPath, "config", "./mercury.yaml", "config file path")
}

var ConfigPath string

func initConfig() {

}

var rootCmd = &cobra.Command{
	Use:     "mercury", //command name
	Short:   "A handwritten distributed instant messaging system Mercury.",
	Long:    `A handwritten distributed instant messaging system. Mercury provides an efficient and secure communication platform designed to enable rapid information exchange and a seamless communication experience. `,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mercury")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
