package cmd

import (
	"context"
	"fmt"

	"github.com/cedws/iapc/iap"
	"github.com/cedws/iapc/internal/proxy"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
)

var (
	zone       string
	ninterface string
)

var instanceCmd = &cobra.Command{
	Use:  "to-instance",
	Long: "Create a tunnel to a remote Compute Engine instance",
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Starting proxy", "dest", fmt.Sprintf("%v:%v", args[0], port), "port", port, "project", project)
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		tokenSource, err := google.DefaultTokenSource(ctx, tokenScopes...)
		if err != nil {
			log.Fatal(err)
		}

		opts := []iap.DialOption{
			iap.WithProject(project),
			iap.WithInstance(args[0], zone, ninterface),
			iap.WithPort(fmt.Sprint(port)),
			iap.WithTokenSource(&tokenSource),
		}
		if compress {
			opts = append(opts, iap.WithCompression())
		}

		proxy.Listen(ctx, listen, opts)
	},
}

func init() {
	instanceCmd.Flags().StringVarP(&zone, "zone", "z", "", "Target zone name")
	instanceCmd.Flags().StringVarP(&ninterface, "interface", "i", "nic0", "Target network interface")
	instanceCmd.MarkFlagRequired("zone")

	rootCmd.AddCommand(instanceCmd)
}
