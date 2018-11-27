package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"wps_store/rpc/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC wps store micro service",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()

		server.Run()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server port")
	rootCmd.AddCommand(serverCmd)
}
