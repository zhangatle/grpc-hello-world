package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"grpc-hello-world/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error : %v\n", err)
			}
		}()

		server.Run()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.Port, "port", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.crt", "cert-pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert-key path")
	serverCmd.Flags().StringVarP(&server.CertServerName, "cert-server-name", "", "localhost", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto", "path to the directory which contains swagger definitions")

	rootCmd.AddCommand(serverCmd)
}