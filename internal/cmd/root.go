package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	logLevel    string
	compress    bool
	listen      string
	project     string
	port        uint
	tokenScopes []string
)

var rootCmd = &cobra.Command{
	Use:  "iapc",
	Long: "Utility for Google Cloud's Identity-Aware Proxy",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			log.Warnf("Could not set log level to %s, use one of: {debug|info|warn|error|fatal}", logLevel)
		}
		log.SetLevel(level)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Set log level")
	rootCmd.PersistentFlags().BoolVarP(&compress, "compress", "c", false, "Enable WebSocket compression")
	rootCmd.PersistentFlags().StringVarP(&listen, "listen", "l", "127.0.0.1:0", "Listen address and port")
	rootCmd.PersistentFlags().StringVar(&project, "project", "", "Project ID")
	rootCmd.PersistentFlags().UintVarP(&port, "port", "p", 22, "Target port")
	rootCmd.PersistentFlags().StringSliceVarP(&tokenScopes, "token-scopes", "s", []string{"https://www.googleapis.com/auth/cloud-platform"}, "Token scopes")
	rootCmd.MarkFlagRequired("project")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
