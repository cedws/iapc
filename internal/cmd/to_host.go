package cmd

import (
	"fmt"

	"github.com/cedws/iapc/iap"
	"github.com/cedws/iapc/internal/proxy"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
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
		log.Info("Started proxy", "listen", listen, "dest", fmt.Sprintf("%v:%v", args[0], port), "project", project)
	},
	Run: func(cmd *cobra.Command, args []string) {
		opts := []iap.DialOption{
			iap.WithProject(project),
			iap.WithHost(args[0], region, network, destGroup),
			iap.WithPort(fmt.Sprint(port)),
			iap.WithTokenSource(getTokenSource()),
		}
		if compress {
			opts = append(opts, iap.WithCompression())
		}

		proxy.Start(listen, opts)
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
