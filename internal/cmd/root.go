package cmd

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	debug       bool
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
		if debug {
			log.SetLevel(log.DebugLevel)
		}
	},
}

func tokenSource() *oauth2.TokenSource {
	tokenSource, err := google.DefaultTokenSource(context.Background(), tokenScopes...)
	if err != nil {
		log.Fatal(err)
	}
	return &tokenSource
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
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
