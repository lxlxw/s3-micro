package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"wps_store/rpc/server_http"
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

		server_http.Run()
	},
}

func init() {
	serverHttpCmd.Flags().StringVarP(&server_http.ServerPort, "server-port", "p", "50052", "server port")
	serverHttpCmd.Flags().StringVarP(&server_http.ServerHttpPort, "http-port", "u", "8088", "http server port")

	rootCmd.AddCommand(serverHttpCmd)
}
