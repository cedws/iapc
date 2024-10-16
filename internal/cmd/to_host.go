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
	destGroup string
	network   string
	region    string
)

var hostCmd = &cobra.Command{
	Use:  "to-host",
	Long: "Create a tunnel to a remote private IP or FQDN (requires BeyondCorp Enterprise)",
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Starting proxy", "dest", fmt.Sprintf("%v:%v", args[0], port), "project", project)
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		tokenSource, err := google.DefaultTokenSource(ctx, tokenScopes...)
		if err != nil {
			log.Fatal(err)
		}

		opts := []iap.DialOption{
			iap.WithProject(project),
			iap.WithHost(args[0], region, network, destGroup),
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
	hostCmd.Flags().StringVarP(&destGroup, "dest-group", "d", "", "Destination group name")
	hostCmd.Flags().StringVarP(&region, "region", "r", "", "Target region name")
	hostCmd.Flags().StringVarP(&network, "network", "n", "", "Target network name")
	hostCmd.MarkFlagRequired("dest-group")
	hostCmd.MarkFlagRequired("region")
	hostCmd.MarkFlagRequired("network")

	rootCmd.AddCommand(hostCmd)
}
