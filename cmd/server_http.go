package cmd

import (
	"log"

	"github.com/lxlxw/s3-micro/api/rpcserver"

	"github.com/spf13/cobra"
)

var serverHttpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run the gRPC wps store http server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()

		rpcserver.RunHttpServer()
	},
}

func init() {
	serverHttpCmd.Flags().StringVarP(&rpcserver.ServerPort, "server-port", "p", "50052", "server port")
	serverHttpCmd.Flags().StringVarP(&rpcserver.ServerHttpPort, "http-port", "u", "8088", "http server port")

	rootCmd.AddCommand(serverHttpCmd)
}
