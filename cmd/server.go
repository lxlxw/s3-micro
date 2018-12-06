package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/lxlxw/s3-micro/api/rpcserver"
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

		rpcserver.RunServer()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&rpcserver.ServerPort, "port", "p", "50052", "server port")
	rootCmd.AddCommand(serverCmd)
}
