package main

import (
	"github.com/rabcat/cn-chess-backend/be"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serverCmd = &cobra.Command{
		Use: "start the backend service",
		Run: func(cmd *cobra.Command, args []string) {
			exec()
		},
	}
)

func init() {
	// bind flags
	serverCmd.Flags().IntP("http_port", "p", 7345, "Port to run http server on")

	viper.BindPFlag("http_port", serverCmd.Flags().Lookup("http_port"))
}

func exec() {
	svc := be.NewCNChessBackendService(viper.GetInt("http_port"))
	err := svc.Start()

	if err != nil {
		panic(err)
	}
}

func main() {
	serverCmd.Execute()
}