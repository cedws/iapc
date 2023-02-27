package cmd

import (
	"fmt"

	"github.com/cedws/iapc/iap"
	"github.com/spf13/cobra"
)

var (
	destGroup string
	network   string
	region    string
)

var hostCmd = &cobra.Command{
	Use:  "host",
	Long: "Create a tunnel to a remote host or FQDN (requires BeyondCorp Enterprise)",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		opts := []iap.DialOption{
			iap.WithProject(project),
			iap.WithHost(args[0], region, network, destGroup),
			iap.WithPort(fmt.Sprint(port)),
		}
		if compress {
			opts = append(opts, iap.WithCompression())
		}

		startProxy(opts)
	},
}

func init() {
	hostCmd.Flags().StringVarP(&zone, "dest-group", "g", "", "Destination group name")
	hostCmd.Flags().StringVarP(&region, "region", "r", "", "Target region name")
	hostCmd.Flags().StringVarP(&network, "network", "e", "", "Target network name")
	hostCmd.MarkFlagRequired("dest-group")
	hostCmd.MarkFlagRequired("region")
	hostCmd.MarkFlagRequired("network")

	rootCmd.AddCommand(hostCmd)
}
