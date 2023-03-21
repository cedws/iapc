package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	compress bool
	listen   string
	project  string
	port     uint
)

var rootCmd = &cobra.Command{
	Use:  "iapc",
	Long: "Utility for Google Cloud's Identity-Aware Proxy",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&compress, "compress", "c", false, "Enable WebSocket compression")
	rootCmd.PersistentFlags().StringVarP(&listen, "listen", "l", "127.0.0.1:0", "Listen address and port")
	rootCmd.PersistentFlags().StringVar(&project, "project", "", "Project ID")
	rootCmd.PersistentFlags().UintVarP(&port, "port", "p", 22, "Target port")
	rootCmd.MarkFlagRequired("project")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
