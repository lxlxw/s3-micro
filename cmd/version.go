package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//TODO: 读取配置文件
var Version = "v1.0"

func ShowVersion() {
	fmt.Printf("wps-store-micro-service version: %s\n", Version)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of wps-store-micro-service",
	Run: func(cmd *cobra.Command, args []string) {
		ShowVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}